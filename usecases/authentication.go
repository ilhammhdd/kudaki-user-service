package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	sarama "gopkg.in/Shopify/sarama.v1"

	"github.com/google/uuid"

	"github.com/golang/protobuf/proto"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/protobuf/ptypes"

	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

func Signup(su *events.SignupRequested, dbOperator DBOperator, esp EventSourceProducer) {
	password, err := bcrypt.GenerateFromPassword([]byte(su.Profile.User.Password), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	su.Profile.User.Password = string(password)

	var eventStatus events.Status
	uves := events.UserVerificationEmailSent{
		Uid:  uuid.New().String(),
		User: su.Profile.User,
	}

	row, err := dbOperator.QueryRow("SELECT count(id) FROM users WHERE email=?", su.Profile.User.Email)
	errorkit.ErrorHandled(err)

	var userID uint
	errorkit.ErrorHandled(row.Scan(&userID))

	if userID > 0 {
		eventStatus.HttpCode = http.StatusConflict
		eventStatus.Errors = []string{"user with the given email already exists"}
		eventStatus.Timestamp = ptypes.TimestampNow()
		uves.EventStatus = &eventStatus
		uvesBytes, err := proto.Marshal(&uves)
		errorkit.ErrorHandled(err)

		esp.Set(entities.Topics_name[int32(entities.Topics_USER_VERIFICATION_EMAIL_SENT)], 0, sarama.OffsetNewest)
		start := time.Now()
		partition, offset, err := esp.SyncProduce(uves.Uid, uvesBytes)
		errorkit.ErrorHandled(err)

		duration := time.Since(start)
		log.Println(duration.Seconds(), " seconds passed after producing")

		log.Println("UserVerificationEmailSent event produced at : ", partition, offset)

		return
	}

	safekit.Do(func() {
		e := &jwtkit.ECDSA{
			PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
			PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

		je := jwtkit.JWTExpiration(172800000)
		jwtString, err := je.GenerateSignedJWTString(
			e,
			"unverified Kudaki.id user",
			"Kudaki.id user service",
			&map[string]interface{}{
				"user_uuid": su.Profile.User.Uuid})
		errorkit.ErrorHandled(err)

		body := fmt.Sprintf("%s/user/verify?verify_token=%s", os.Getenv("GATEWAY_HOST"), string(jwtString))

		mail := Mail{
			From: mail.Address{
				Name:    "Notification Kudaki.id",
				Address: os.Getenv("MAIL")},
			To: mail.Address{
				Name:    su.Profile.FullName,
				Address: su.Profile.User.Email},
			Subject: "User account verification",
			Body:    []byte(body)}

		if errorkit.ErrorHandled(mail.SendWithTLS()) {
			eventStatus.HttpCode = http.StatusBadRequest
			eventStatus.Errors = []string{"error occured when sending verification email"}
			eventStatus.Timestamp = ptypes.TimestampNow()
			uves.EventStatus = &eventStatus
			uvesBytes, err := proto.Marshal(&uves)
			errorkit.ErrorHandled(err)

			esp.Set(entities.Topics_name[int32(entities.Topics_USER_VERIFICATION_EMAIL_SENT)], 0, sarama.OffsetNewest)
			start := time.Now()
			partition, offset, err := esp.SyncProduce(uves.Uid, uvesBytes)
			errorkit.ErrorHandled(err)

			duration := time.Since(start)
			log.Println(duration.Seconds(), " seconds passed after producing")

			log.Println("UserVerificationEmailSent event produced at : ", partition, offset)
		}
	})

	createUserAndProfile(su, dbOperator)

	eventStatus.HttpCode = http.StatusOK
	eventStatus.Messages = []string{"successfully sent verfication email"}
	eventStatus.Timestamp = ptypes.TimestampNow()
	uves.EventStatus = &eventStatus
	uvesBytes, err := proto.Marshal(&uves)
	errorkit.ErrorHandled(err)

	log.Println("UserVerificationEmailSent", uves.Uid)

	esp.Set(entities.Topics_name[int32(entities.Topics_USER_VERIFICATION_EMAIL_SENT)], 0, sarama.OffsetNewest)
	start := time.Now()
	_, _, err = esp.SyncProduce(uves.Uid, uvesBytes)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Println(duration.Seconds(), " seconds passed after producing")
}

func VerifyUser(vu *events.VerifyUserRequested, dbOperator DBOperator, esp EventSourceProducer) {
	e := jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	log.Println("verify jwt: ", jwtkit.JWTString(vu.VerifyUserJwt))

	verified, err := jwtkit.VerifyJWTString(&e, jwtkit.JWTString(vu.VerifyUserJwt))
	errorkit.ErrorHandled(err)

	sdu := events.Signedup{
		Uid: vu.Uid,
		EventStatus: &events.Status{
			Timestamp: ptypes.TimestampNow()}}

	if verified {
		validated, err := jwtkit.ValidateExpired(jwtkit.JWTString(vu.VerifyUserJwt))
		errorkit.ErrorHandled(err)

		if validated {
			jwt, err := jwtkit.GetJWT(jwtkit.JWTString(vu.VerifyUserJwt))
			errorkit.ErrorHandled(err)

			err = dbOperator.Command("DELETE FROM unverified_users WHERE user_uuid=?", jwt.Payload.Claims["user_uuid"])
			errorkit.ErrorHandled(err)

			sdu.EventStatus.HttpCode = http.StatusOK
			sdu.EventStatus.Messages = append(sdu.EventStatus.Messages, "user account verified")
		} else {
			sdu.EventStatus.HttpCode = http.StatusUnauthorized
			sdu.EventStatus.Errors = append(sdu.EventStatus.Messages, errors.New("jwt expired").Error())
		}
	} else {
		sdu.EventStatus.HttpCode = http.StatusUnauthorized
		sdu.EventStatus.Errors = append(sdu.EventStatus.Messages, errors.New("user unverified").Error())
	}

	sduBytes, err := proto.Marshal(&sdu)
	errorkit.ErrorHandled(err)

	esp.Set(entities.Topics_name[int32(entities.Topics_SIGNED_UP)], 0, sarama.OffsetNewest)
	esp.SyncProduce(sdu.Uid, sduBytes)
}

func Login(lr *events.LoginRequested, dbo DBOperator, esp EventSourceProducer) {
	log.Println("lr :", lr)
	row, err := dbo.QueryRow("SELECT * FROM users WHERE email=?", lr.User.Email)
	errorkit.ErrorHandled(err)

	var usr user.User
	var usrID uint64
	var role string
	var accountType string

	err = row.Scan(&usrID, &usr.Uuid, &usr.Email, &usr.Password, &usr.Token, &role, &usr.PhoneNumber, &accountType)

	if err == sql.ErrNoRows {
		err = produceLoggedin(esp, lr.Uid, []string{"user with the given email doesn't exists"}, nil, http.StatusUnauthorized, &usr)
		if !errorkit.ErrorHandled(err) {
			return
		}
	}

	row, err = dbo.QueryRow("SELECT id FROM unverified_users WHERE user_uuid=?", usr.Uuid)
	var unverifiedID uint64
	err = row.Scan(&unverifiedID)

	log.Println("the unverified user id :", unverifiedID)

	if err != sql.ErrNoRows {
		err = produceLoggedin(esp, lr.Uid, []string{"user unverified"}, nil, http.StatusUnauthorized, &usr)
		if !errorkit.ErrorHandled(err) {
			return
		}
	}

	usr.Role = user.Role(user.Role_value[role])
	usr.AccountType = user.AccountType(user.AccountType_value[accountType])

	log.Println("user in LoginRequested event : ", usr)

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(lr.User.Password)); err != nil {
		err = produceLoggedin(esp, lr.Uid, []string{"wrong password"}, nil, http.StatusUnauthorized, &usr)
		if !errorkit.ErrorHandled(err) {
			return
		}
	}

	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	jwtString, err := jwtkit.JWTExpiration(5.256e+9).GenerateSignedJWTString(
		e,
		"verified Kudaki.id user",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user_uuid": usr.Uuid})

	errorkit.ErrorHandled(err)

	err = dbo.Command("UPDATE users SET token=? WHERE uuid=?", string(jwtString), usr.Uuid)
	errorkit.ErrorHandled(err)

	usr.Token = string(jwtString)
	produceLoggedin(esp, lr.Uid, nil, []string{"successfully logged in"}, http.StatusOK, &usr)
}

func ResetPassword(rpr *events.ResetPasswordRequested, dbo DBOperator, esp EventSourceProducer) {

	var oldPasswordHashed []byte

	row, err := dbo.QueryRow("SELECT password FROM users WHERE uuid=?", rpr.Profile.User.Uuid)
	errorkit.ErrorHandled(err)
	err = row.Scan(&oldPasswordHashed)
	errorkit.ErrorHandled(err)

	if compareErr := bcrypt.CompareHashAndPassword(oldPasswordHashed, []byte(rpr.OldPassword)); compareErr != nil {
		producePasswordReseted(esp, []string{compareErr.Error()}, http.StatusUnauthorized, rpr.Uid)
		return
	}

	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(rpr.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	log.Println("rpr in ResetPassword usecase :", rpr)

	err = dbo.Command("UPDATE users SET password=? WHERE uuid=?", string(newPasswordHashed), rpr.Profile.User.Uuid)
	if !errorkit.ErrorHandled(err) {
		mail := Mail{
			Body: []byte("Your password has changed"),
			From: mail.Address{
				Address: os.Getenv("MAIL"),
				Name:    "Notification kudaki.id"},
			Subject: "Password changed",
			To: mail.Address{
				Address: rpr.Profile.User.Email,
				Name:    rpr.Profile.FullName}}

		errMail := mail.SendWithTLS()

		if !errorkit.ErrorHandled(errMail) {
			log.Println("successfully changed password and send email")
			producePasswordReseted(esp, nil, int32(http.StatusOK), rpr.Uid)
		} else {
			producePasswordReseted(esp, []string{errMail.Error()}, int32(http.StatusInternalServerError), rpr.Uid)
		}
	} else {
		producePasswordReseted(esp, []string{err.Error()}, http.StatusInternalServerError, rpr.Uid)
	}

}

func producePasswordReseted(esp EventSourceProducer, errs []string, httpCode int32, uid string) {
	esp.Set(events.User_name[int32(events.User_PASSWORD_RESETED)], 0, sarama.OffsetNewest)

	pr := &events.PasswordReseted{
		EventStatus: &events.Status{
			Errors:    errs,
			HttpCode:  httpCode,
			Timestamp: ptypes.TimestampNow()},
		Uid: uid,
	}
	prBytes, err := proto.Marshal(pr)
	errorkit.ErrorHandled(err)

	_, _, err = esp.SyncProduce(uid, prBytes)
	errorkit.ErrorHandled(err)
}

func produceLoggedin(esp EventSourceProducer, requestUID string, errs []string, msg []string, httpCode int32, usr *user.User) error {
	esp.Set(events.User_name[int32(events.User_LOGGED_IN)], 0, sarama.OffsetNewest)

	loggedin := events.Loggedin{
		Uid:         requestUID,
		EventStatus: new(events.Status)}

	loggedin.EventStatus.Errors = errs
	loggedin.EventStatus.Messages = msg
	loggedin.EventStatus.HttpCode = httpCode
	loggedin.EventStatus.Timestamp = ptypes.TimestampNow()
	loggedin.User = usr
	loggedinBytes, err := proto.Marshal(&loggedin)
	errorkit.ErrorHandled(err)

	_, _, err = esp.SyncProduce(loggedin.Uid, loggedinBytes)

	return err
}

func createUserAndProfile(su *events.SignupRequested, dbo DBOperator) {
	dbo.Command(
		"INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES(?,?,?,?,?,?,?)",
		su.Profile.User.Uuid,
		su.Profile.User.Email,
		su.Profile.User.Password,
		"",
		user.Role_name[int32(su.Profile.User.Role)],
		su.Profile.User.PhoneNumber,
		user.AccountType_name[int32(su.Profile.User.AccountType)],
	)

	dbo.Command(
		"INSERT INTO unverified_users(user_uuid) VALUES(?)",
		su.Profile.User.Uuid,
	)

	dbo.Command("INSERT INTO profiles(user_uuid,uuid,full_name,reputation) VALUES(?,?,?,?)", su.Profile.User.Uuid, uuid.New().String(), su.Profile.FullName, 0)
}
