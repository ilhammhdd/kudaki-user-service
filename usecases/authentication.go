package usecases

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"

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

func (s *Signup) sendVerifEmail(usr *user.User, profile *user.Profile) (verifyToken string, mailErr error) {
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

func (s *Signup) produceVerifEmailSent(usr *user.User, verifyToken string) {
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

func (s *Signup) insertUserAndProfile(usr *user.User, profile *user.Profile) error {
	_, err := s.DBO.Command("INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES (?,?,?,?,?,?,?);",
		usr.Uuid, usr.Email, usr.Password, usr.Token, usr.Role.String(), usr.PhoneNumber, usr.AccountType.String())
	if errorkit.ErrorHandled(err) {
		return err
	}

	_, err = s.DBO.Command("INSERT INTO profiles(uuid,user_uuid,full_name,photo,reputation) VALUES (?,?,?,?,?);",
		profile.Uuid, profile.User.Uuid, profile.FullName, profile.Photo, profile.Reputation)
	if errorkit.ErrorHandled(err) {
		return err
	}
	return nil
}

func (s *Signup) indexUserAndProfile(usr *user.User, profile *user.Profile) error {
	userClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	userClient.CreateIndex(kudakiredisearch.User.Schema())

	userDoc := redisearch.NewDocument(kudakiredisearch.RedisearchText(usr.Uuid).Sanitize(), 1.0)
	userDoc.Set("user_uuid", usr.Uuid)
	userDoc.Set("user_email", usr.Email)
	userDoc.Set("user_password", usr.Password)
	userDoc.Set("user_token", usr.Token)
	userDoc.Set("user_role", usr.Role.String())
	userDoc.Set("user_phone_number", usr.PhoneNumber)
	userDoc.Set("user_account_type", usr.AccountType.String())

	err := userClient.IndexOptions(redisearch.DefaultIndexingOptions, userDoc)
	if errorkit.ErrorHandled(err) {
		return err
	}

	profileClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Profile.Name())
	profileClient.CreateIndex(kudakiredisearch.Profile.Schema())

	profileDoc := redisearch.NewDocument(kudakiredisearch.RedisearchText(profile.Uuid).Sanitize(), 1.0)
	profileDoc.Set("profile_uuid", profile.Uuid)
	profileDoc.Set("profile_full_name", profile.FullName)
	profileDoc.Set("profile_photo", profile.Photo)
	profileDoc.Set("profile_reputation", profile.Reputation)

	err = profileClient.IndexOptions(redisearch.DefaultIndexingOptions, profileDoc)
	if errorkit.ErrorHandled(err) {
		return err
	}

	return nil
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
	if verifyToken, mailErr := s.sendVerifEmail(newUser, newProfile); mailErr != nil {
		outEvent.EventStatus.HttpCode = http.StatusInternalServerError
		outEvent.EventStatus.Errors = []string{"error occured while sending verification email"}
		return outEvent
	} else {
		s.produceVerifEmailSent(newUser, verifyToken)
	}

	if s.insertUserAndProfile(newUser, newProfile) != nil {
		outEvent.EventStatus.HttpCode = http.StatusInternalServerError
		outEvent.EventStatus.Errors = []string{"error occured while inserting new user and profile"}
		return outEvent
	}

	if s.indexUserAndProfile(newUser, newProfile) != nil {
		outEvent.EventStatus.HttpCode = http.StatusInternalServerError
		outEvent.EventStatus.Errors = []string{"error occured while indexing new user and profile"}
		return outEvent
	}

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Profile = newProfile

	return outEvent
}
