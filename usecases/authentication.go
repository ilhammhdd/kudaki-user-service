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
