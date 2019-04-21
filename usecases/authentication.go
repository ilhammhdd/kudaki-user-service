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

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/golang/protobuf/proto"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/protobuf/ptypes"

	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

func Signup(su *events.SignupRequested, dbOperator DBOperator, esp EventDrivenProducer) {
	password, err := bcrypt.GenerateFromPassword([]byte(su.Profile.User.Password), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	su.Profile.User.Password = string(password)

	var eventStatus events.Status
	sdu := events.Signedup{
		Uid:  su.Uid,
		User: su.Profile.User,
	}
	row, err := dbOperator.QueryRow("SELECT count(id) FROM users WHERE email=?", su.Profile.User.Email)
	errorkit.ErrorHandled(err)

	var totalID uint
	err = row.Scan(&totalID)

	if totalID > 0 {
		eventStatus.HttpCode = http.StatusConflict
		eventStatus.Errors = []string{"user with the given email already exists"}
		eventStatus.Timestamp = ptypes.TimestampNow()
		sdu.EventStatus = &eventStatus
		uvesBytes, err := proto.Marshal(&sdu)
		errorkit.ErrorHandled(err)

		esp.Set(entities.Topics_name[int32(entities.Topics_SIGNED_UP)])
		start := time.Now()
		partition, offset, err := esp.SyncProduce(sdu.Uid, uvesBytes)
		errorkit.ErrorHandled(err)

		duration := time.Since(start)
		log.Printf("produced Signedup : partition = %d, offset = %d, duration = %f seconds", partition, offset, duration.Seconds())
		return
	}

	safekit.Do(func() {
		var uves events.UserVerificationEmailSent
		uves.Uid = su.Uid
		uves.User = su.Profile.User

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

			esp.Set(entities.Topics_name[int32(entities.Topics_USER_VERIFICATION_EMAIL_SENT)])
			start := time.Now()
			partition, offset, err := esp.SyncProduce(uves.Uid, uvesBytes)
			errorkit.ErrorHandled(err)

			duration := time.Since(start)
			log.Printf("produced UserEmailVerificationSent : partition = %d, offset = %d, time = %f seconds", partition, offset, duration.Seconds())
		}
	})

	createUserAndProfile(su, dbOperator)

	eventStatus.HttpCode = http.StatusOK
	eventStatus.Messages = []string{"please verify your account by clicking the link we sent to your email"}
	eventStatus.Timestamp = ptypes.TimestampNow()
	sdu.EventStatus = &eventStatus
	uvesBytes, err := proto.Marshal(&sdu)
	errorkit.ErrorHandled(err)

	esp.Set(entities.Topics_name[int32(entities.Topics_SIGNED_UP)])
	start := time.Now()
	partition, offset, err := esp.SyncProduce(sdu.Uid, uvesBytes)
	errorkit.ErrorHandled(err)
	duration := time.Since(start)
	log.Printf("produced Signedup : partition = %d, offset = %d, time = %f seconds", partition, offset, duration.Seconds())
}

func VerifyUser(vu *events.VerifyUserRequested, dbOperator DBOperator) *events.Signedup {
	e := jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	signedUp := events.Signedup{
		EventStatus: new(events.Status),
		Uid:         vu.Uid,
	}

	verified, err := jwtkit.VerifyJWTString(&e, jwtkit.JWTString(vu.VerifyUserJwt))
	errorkit.ErrorHandled(err)

	if verified {
		validated, err := jwtkit.ValidateExpired(jwtkit.JWTString(vu.VerifyUserJwt))
		errorkit.ErrorHandled(err)

		if validated {
			jwt, err := jwtkit.GetJWT(jwtkit.JWTString(vu.VerifyUserJwt))
			errorkit.ErrorHandled(err)

			err = dbOperator.Command("DELETE FROM unverified_users WHERE user_uuid=?", jwt.Payload.Claims["user_uuid"])
			errorkit.ErrorHandled(err)

			signedUp.EventStatus.HttpCode = http.StatusOK
			signedUp.EventStatus.Messages = []string{"user account verified"}
		} else {
			signedUp.EventStatus.HttpCode = http.StatusUnauthorized
			signedUp.EventStatus.Errors = []string{errors.New("jwt expired").Error()}
		}
	} else {
		signedUp.EventStatus.HttpCode = http.StatusUnauthorized
		signedUp.EventStatus.Errors = []string{errors.New("user unverified").Error()}
	}

	return &signedUp
}

func Login(lr *events.LoginRequested, dbo DBOperator) *events.Loggedin {
	var usr user.User
	var usrID uint64
	var role string
	var accountType string
	loggedIn := events.Loggedin{
		EventStatus: new(events.Status),
		Uid:         lr.Uid}

	row, err := dbo.QueryRow("SELECT * FROM users WHERE email=?", lr.User.Email)
	errorkit.ErrorHandled(err)

	err = row.Scan(&usrID, &usr.Uuid, &usr.Email, &usr.Password, &usr.Token, &role, &usr.PhoneNumber, &accountType)

	if err == sql.ErrNoRows {
		loggedIn.EventStatus.Errors = []string{"user with the given email doesn't exists"}
		loggedIn.EventStatus.HttpCode = http.StatusUnauthorized
		loggedIn.User = &usr

		return &loggedIn
	}

	row, err = dbo.QueryRow("SELECT id FROM unverified_users WHERE user_uuid=?", usr.Uuid)
	var unverifiedID uint64
	err = row.Scan(&unverifiedID)

	if err != sql.ErrNoRows {
		loggedIn.EventStatus.Errors = []string{"user unverified"}
		loggedIn.EventStatus.HttpCode = http.StatusUnauthorized
		loggedIn.User = &usr

		return &loggedIn
	}

	usr.Role = user.Role(user.Role_value[role])
	usr.AccountType = user.AccountType(user.AccountType_value[accountType])

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(lr.User.Password)); err != nil {
		loggedIn.EventStatus.Errors = []string{"wrong password"}
		loggedIn.EventStatus.HttpCode = http.StatusUnauthorized
		loggedIn.User = &usr

		return &loggedIn
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

	loggedIn.EventStatus.Errors = []string{"successfully logged in"}
	loggedIn.EventStatus.HttpCode = http.StatusOK
	loggedIn.User = &usr

	return &loggedIn
}

func ResetPassword(rpr *events.ResetPasswordRequested, dbo DBOperator) *events.PasswordReseted {

	var oldPasswordHashed []byte
	passwordReseted := events.PasswordReseted{
		EventStatus: new(events.Status),
		Uid:         rpr.Uid}

	row, err := dbo.QueryRow("SELECT password FROM users WHERE uuid=?", rpr.Profile.User.Uuid)
	errorkit.ErrorHandled(err)
	err = row.Scan(&oldPasswordHashed)
	errorkit.ErrorHandled(err)

	if compareErr := bcrypt.CompareHashAndPassword(oldPasswordHashed, []byte(rpr.OldPassword)); compareErr != nil {
		passwordReseted.EventStatus.Errors = []string{compareErr.Error()}
		passwordReseted.EventStatus.HttpCode = http.StatusUnauthorized

		return &passwordReseted
	}

	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(rpr.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	dbo.Command("UPDATE users SET password=? WHERE uuid=?", string(newPasswordHashed), rpr.Profile.User.Uuid)

	safekit.Do(func() {
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
		if errorkit.ErrorHandled(errMail) {
			log.Println("failed to send password changed email")
		} else {
			log.Println("successfully send password changed email")
		}
	})

	passwordReseted.EventStatus.HttpCode = http.StatusOK

	return &passwordReseted
}

func producePasswordReseted(esp EventDrivenProducer, errs []string, httpCode int32, uid string) {
	esp.Set(events.User_name[int32(events.User_PASSWORD_RESETED)])

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

func produceLoggedin(esp EventDrivenProducer, requestUID string, errs []string, msg []string, httpCode int32, usr *user.User) error {
	esp.Set(events.User_name[int32(events.User_LOGGED_IN)])

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
