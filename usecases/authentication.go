package usecases

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-entities/user"
	"golang.org/x/crypto/bcrypt"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/golang/protobuf/proto"
)

type Signup struct {
	DBO      DBOperator
	Producer EventDrivenProducer
}

func (s *Signup) initInOutEvent(in proto.Message) (inEvent *events.SignupRequested, outEvent *events.Signedup) {
	inEvent = in.(*events.SignupRequested)

	outEvent = new(events.Signedup)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.SignupRequested = inEvent
	outEvent.Uid = inEvent.Uid

	return inEvent, outEvent
}

func (s *Signup) userExists(inEvent *events.SignupRequested) bool {
	row, err := s.DBO.QueryRow("SELECT id FROM users WHERE email = ?;", inEvent.Email)
	errorkit.ErrorHandled(err)

	var existedUserID uint64

	if row.Scan(&existedUserID) == sql.ErrNoRows {
		return false
	}

	return true
}

func (s *Signup) initUserAndProfile(inEvent *events.SignupRequested) (*user.User, *user.Profile) {
	usr := new(user.User)
	usr.AccountType = user.AccountType_NATIVE
	usr.Email = inEvent.Email
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(inEvent.Password), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	usr.Password = string(encryptedPassword)
	usr.PhoneNumber = inEvent.PhoneNumber
	usr.Role = user.Role(user.Role_value[inEvent.Role])
	usr.Uuid = uuid.New().String()

	profile := new(user.Profile)
	profile.FullName = inEvent.FullName
	profile.Photo = inEvent.Photo
	profile.Reputation = 0
	profile.User = usr
	profile.Uuid = uuid.New().String()

	return usr, profile
}

func (s *Signup) sendVerifyEmail(usr *user.User, profile *user.Profile) (verifyToken string, mailErr error) {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	je := jwtkit.JWTExpiration(172800000)
	jwtString, err := je.GenerateSignedJWTString(
		e,
		"unverified Kudaki.id user",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user_uuid": usr.Uuid})
	errorkit.ErrorHandled(err)

	body := fmt.Sprintf("%s/user/verify?verify_token=%s", os.Getenv("GATEWAY_HOST"), string(jwtString))

	mail := Mail{
		From: mail.Address{
			Name:    "Notification Kudaki.id",
			Address: os.Getenv("MAIL")},
		To: mail.Address{
			Name:    profile.FullName,
			Address: usr.Email},
		Subject: "User account verification",
		Body:    []byte(body)}

	mailErr = mail.SendWithTLS()
	if errorkit.ErrorHandled(mailErr) {
		return "", mailErr
	}
	return string(jwtString), nil
}

func (s *Signup) produceVerifyEmailSent(usr *user.User, verifyToken string) {
	uves := new(events.UserVerificationEmailSent)
	uves.EventStatus = new(events.Status)
	uves.Uid = uuid.New().String()
	uves.User = usr
	uves.VerifyToken = verifyToken

	uvesByte, err := proto.Marshal(uves)
	errorkit.ErrorHandled(err)

	s.Producer.Set(events.UserTopic_USER_VERIFICATION_EMAIL_SENT.String())
	start := time.Now()
	partition, offset, err := s.Producer.SyncProduce(uves.Uid, uvesByte)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)

	log.Printf("produced %s : partition = %d, offset = %d, key = %s, duration = %f seconds",
		events.UserTopic_USER_VERIFICATION_EMAIL_SENT.String(), partition, offset, uves.Uid, duration.Seconds())
}

func (s *Signup) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := s.initInOutEvent(in)

	ok := s.userExists(inEvent)
	if ok {
		outEvent.EventStatus.HttpCode = http.StatusBadRequest
		outEvent.EventStatus.Errors = []string{"user with the given email already exists"}
		return outEvent
	}

	newUser, newProfile := s.initUserAndProfile(inEvent)
	if verifyToken, mailErr := s.sendVerifyEmail(newUser, newProfile); mailErr != nil {
		outEvent.EventStatus.HttpCode = http.StatusInternalServerError
		outEvent.EventStatus.Errors = []string{"error occured while sending verification email"}
		return outEvent
	} else {
		s.produceVerifyEmailSent(newUser, verifyToken)
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Profile = newProfile

	return outEvent
}

type Login struct {
	DBO DBOperator
}

func (l *Login) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := l.initInOutEvent(in)

	l.initInOutEvent(in)
	existedUser, ok := l.userExists(inEvent)
	if !ok {
		outEvent.EventStatus.Errors = []string{"user with the given email doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if !l.userVerified(existedUser.Uuid) {
		outEvent.EventStatus.Errors = []string{"user wasn't verified"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	existedProfile, ok := l.profileExists(inEvent, existedUser)
	if !ok {
		outEvent.EventStatus.Errors = []string{"profile corresponds to user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(inEvent.Password)); err != nil {
		outEvent.EventStatus.Errors = []string{"wrong password"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	existedUser.Token = l.generateAuthToken(existedUser, existedProfile)
	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = existedUser

	return outEvent
}

func (l *Login) initInOutEvent(in proto.Message) (inEvent *events.LoginRequested, outEvent *events.Loggedin) {
	inEvent = in.(*events.LoginRequested)

	outEvent = new(events.Loggedin)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.LoginRequested = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (l *Login) userExists(inEvent *events.LoginRequested) (*user.User, bool) {
	row, err := l.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM users WHERE email = ?;", inEvent.Email)
	errorkit.ErrorHandled(err)

	var existedUser user.User
	var role, accountType string
	if row.Scan(&existedUser.Uuid, &existedUser.Email, &existedUser.Password, &existedUser.Token, &role, &existedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	existedUser.Role = user.Role(user.Role_value[role])
	existedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &existedUser, true
}

func (l *Login) profileExists(inEvent *events.LoginRequested, existedUser *user.User) (*user.Profile, bool) {
	row, err := l.DBO.QueryRow("SELECT uuid,full_name,photo,reputation FROM profiles WHERE user_uuid = ?;", existedUser.Uuid)
	errorkit.ErrorHandled(err)

	var existedProfile user.Profile
	if row.Scan(&existedProfile.Uuid, &existedProfile.FullName, &existedProfile.Photo, &existedProfile.Reputation) == sql.ErrNoRows {
		return nil, false
	}
	existedProfile.User = existedUser

	return &existedProfile, true
}

func (l *Login) userVerified(userUUID string) bool {
	row, err := l.DBO.QueryRow("SELECT id FROM unverified_users WHERE user_uuid = ?;", userUUID)
	errorkit.ErrorHandled(err)

	var unverifiedUserID uint64
	if row.Scan(&unverifiedUserID) == sql.ErrNoRows {
		return true
	}
	return false
}

func (l *Login) generateAuthToken(usr *user.User, profile *user.Profile) string {
	authECDSA := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	jwtString, err := jwtkit.JWTExpiration(5.256e+9).GenerateSignedJWTString(
		authECDSA,
		"verified Kudaki.id user",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user": map[string]interface{}{
				"account_type": usr.AccountType.String(),
				"email":        usr.Email,
				"phone_number": usr.PhoneNumber,
				"role":         usr.Role.String(),
				"uuid":         usr.Uuid,
			},
			"profile": map[string]interface{}{
				"user_uuid":  profile.User.Uuid,
				"uuid":       profile.Uuid,
				"full_name":  profile.FullName,
				"photo":      profile.Photo,
				"reputation": profile.Reputation,
			},
		})
	errorkit.ErrorHandled(err)

	return string(jwtString)
}

type VerifyUser struct {
	DBO DBOperator
}

func (vu *VerifyUser) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := vu.initInOutEvent(in)

	ecdsaPair := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	if !vu.verifyToken(ecdsaPair, inEvent) {
		outEvent.EventStatus.Errors = []string{"token not verified"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	if !vu.validateToken(inEvent) {
		outEvent.EventStatus.Errors = []string{"token not valid"}
		outEvent.EventStatus.HttpCode = http.StatusUnauthorized
		return outEvent
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.User = vu.retrieveUser(inEvent)

	return outEvent
}

func (vu *VerifyUser) initInOutEvent(in proto.Message) (inEvent *events.VerifyUserRequested, outEvent *events.UserVerified) {
	inEvent = in.(*events.VerifyUserRequested)

	outEvent = new(events.UserVerified)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.VerifyUserRequested = inEvent
	return
}

func (vu *VerifyUser) verifyToken(ecdsaPair *jwtkit.ECDSA, inEvent *events.VerifyUserRequested) bool {
	verified, err := jwtkit.VerifyJWTString(ecdsaPair, jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)
	return verified
}

func (vu *VerifyUser) validateToken(inEvent *events.VerifyUserRequested) bool {
	validated, err := jwtkit.ValidateExpired(jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)
	return validated
}

func (vu *VerifyUser) retrieveUser(inEvent *events.VerifyUserRequested) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(inEvent.VerifyToken))
	errorkit.ErrorHandled(err)

	userUUID := jwt.Payload.Claims["user_uuid"].(string)
	row, err := vu.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM users WHERE uuid=?;", userUUID)
	errorkit.ErrorHandled(err)

	var usr user.User
	var role string
	var accountType string
	err = row.Scan(&usr.Uuid, &usr.Email, &usr.Password, &usr.Token, &role, &usr.PhoneNumber, &accountType)
	errorkit.ErrorHandled(err)

	usr.Role = user.Role(user.Role_value[role])
	usr.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &usr
}

type ChangePassword struct {
	DBO DBOperator
}

func (cp *ChangePassword) initInOutEvent(in proto.Message) (inEvent *events.ChangePasswordRequested, outEvent *events.PasswordChanged) {
	inEvent = in.(*events.ChangePasswordRequested)

	outEvent = new(events.PasswordChanged)
	outEvent.ChangePasswordRequested = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (cp *ChangePassword) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := cp.initInOutEvent(in)
	log.Println(inEvent, outEvent)

	userFromJWT := cp.getUserFromKudakiToken(inEvent.KudakiToken)
	existedUser, ok := cp.userExists(userFromJWT)

	if !ok {
		outEvent.EventStatus.Errors = []string{"user doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	if bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(inEvent.OldPassword)) != nil {
		outEvent.EventStatus.Errors = []string{"wrong old password"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(inEvent.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	existedUser.Password = string(newPassword)

	outEvent.User = existedUser
	outEvent.EventStatus.HttpCode = http.StatusOK
	return outEvent
}

func (cp *ChangePassword) userExists(usr *user.User) (*user.User, bool) {
	row, err := cp.DBO.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM users WHERE uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var retrievedUser user.User
	var role string
	var accountType string
	if row.Scan(&retrievedUser.Uuid, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Token, &role, &retrievedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	retrievedUser.Role = user.Role(user.Role_value[role])
	retrievedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &retrievedUser, true
}

func (cp *ChangePassword) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}
