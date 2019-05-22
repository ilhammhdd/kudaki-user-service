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

		esp.Set(events.UserTopic_name[int32(events.UserTopic_SIGNED_UP)])
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

			esp.Set(events.UserTopic_name[int32(events.UserTopic_USER_VERIFICATION_EMAIL_SENT)])
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

	esp.Set(events.UserTopic_name[int32(events.UserTopic_SIGNED_UP)])
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

	var profile = user.Profile{
		User: new(user.User),
	}
	row, err = dbo.QueryRow("SELECT user_uuid, uuid, full_name, photo, reputation FROM profiles WHERE user_uuid=?", usr.Uuid)
	errorkit.ErrorHandled(err)
	err = row.Scan(&profile.User.Uuid, &profile.Uuid, &profile.FullName, &profile.Photo, &profile.Reputation)
	errorkit.ErrorHandled(err)

	jwtString, err := jwtkit.JWTExpiration(5.256e+9).GenerateSignedJWTString(
		e,
		"verified Kudaki.id user",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user": map[string]interface{}{
				"account_type": accountType,
				"email":        usr.Email,
				"phone_number": usr.PhoneNumber,
				"role":         role,
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

	err = dbo.Command("UPDATE users SET token=? WHERE uuid=?", string(jwtString), usr.Uuid)
	errorkit.ErrorHandled(err)

	usr.Token = string(jwtString)

	loggedIn.EventStatus.Errors = []string{"successfully logged in"}
	loggedIn.EventStatus.HttpCode = http.StatusOK
	loggedIn.User = &usr

	return &loggedIn
}

func ChangePassword(rpr *events.ChangePasswordRequested, dbo DBOperator) *events.PasswordChanged {

	var oldPasswordHashed []byte
	passwordChanged := events.PasswordChanged{
		EventStatus: new(events.Status),
		Uid:         rpr.Uid}

	row, err := dbo.QueryRow("SELECT password FROM users WHERE email=?", rpr.Profile.User.Email)
	errorkit.ErrorHandled(err)
	err = row.Scan(&oldPasswordHashed)
	errorkit.ErrorHandled(err)

	if compareErr := bcrypt.CompareHashAndPassword(oldPasswordHashed, []byte(rpr.OldPassword)); compareErr != nil {
		passwordChanged.EventStatus.Errors = []string{compareErr.Error()}
		passwordChanged.EventStatus.HttpCode = http.StatusUnauthorized

		return &passwordChanged
	}

	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(rpr.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	dbo.Command("UPDATE users SET password=? WHERE email=?", string(newPasswordHashed), rpr.Profile.User.Email)

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

	passwordChanged.EventStatus.HttpCode = http.StatusOK

	return &passwordChanged
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

	dbo.Command("INSERT INTO profiles(user_uuid,uuid,full_name,photo,reputation) VALUES(?,?,?,?,?)",
		su.Profile.User.Uuid,
		uuid.New().String(),
		su.Profile.FullName,
		"",
		0)
}

type SendResetPasswordEmail struct {
	In  *events.SendResetPasswordEmailRequested
	DBO DBOperator
}

func (pr SendResetPasswordEmail) SendEmail() *events.ResetPasswordEmailSent {

	var rpes events.ResetPasswordEmailSent
	rpes.Uid = pr.In.Uid
	rpes.EventStatus = &events.Status{}

	// check if user exists
	row, err := pr.DBO.QueryRow("SELECT u.uuid,u.email,p.full_name FROM users u JOIN profiles p ON p.user_uuid=u.uuid WHERE email = ?;", pr.In.Email)
	errorkit.ErrorHandled(err)

	var user user.User
	var fullName string
	if row.Scan(&user.Uuid, &user.Email, &fullName) == sql.ErrNoRows {
		rpes.EventStatus.Errors = []string{"user with the given email not exists"}
		rpes.EventStatus.HttpCode = http.StatusBadRequest
		rpes.EventStatus.Timestamp = ptypes.TimestampNow()

		return &rpes
	}
	// check if user exists

	// generate reset token
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY")}

	je := jwtkit.JWTExpiration(86400000)
	resetJWT, err := je.GenerateSignedJWTString(
		e,
		"Kudaki.id user resetting password",
		"Kudaki.id user service",
		&map[string]interface{}{
			"user_uuid": user.Uuid,
			"full_name": fullName,
		})
	errorkit.ErrorHandled(err)
	// generate reset token

	// check reset password, insert if not exists update if does
	row, err = pr.DBO.QueryRow("SELECT id FROM reset_passwords WHERE user_uuid = ?;", user.Uuid)
	errorkit.ErrorHandled(err)

	var resetPasswordID uint64

	if row.Scan(&resetPasswordID) == sql.ErrNoRows {
		err = pr.DBO.Command("INSERT INTO reset_passwords(user_uuid,token) VALUES(?,?);", user.Uuid, string(resetJWT))
		errorkit.ErrorHandled(err)
	} else {
		err = pr.DBO.Command("UPDATE reset_passwords SET token=? WHERE user_uuid=?;", string(resetJWT), user.Uuid)
		errorkit.ErrorHandled(err)
	}
	// check reset password, insert if not exists update if does

	// send email contains link to reset password
	// safekit.Do(func() {
	getResetPasswordPageLink := fmt.Sprintf("%s/user/reset-password?reset_token=%s", os.Getenv("GATEWAY_HOST"), resetJWT)

	mail := Mail{
		From: mail.Address{
			Address: os.Getenv("MAIL"),
			Name:    "Kudaki.id account management",
		},
		Body:    []byte(getResetPasswordPageLink),
		Subject: "Reset password user",
		To: mail.Address{
			Address: user.Email,
			Name:    fullName,
		},
	}

	err = mail.SendWithTLS()
	if !errorkit.ErrorHandled(err) {
		rpes.EventStatus.HttpCode = http.StatusOK
		rpes.EventStatus.Timestamp = ptypes.TimestampNow()
	}
	// })
	// send email contains link to reset password

	return &rpes
}

type ResetPassword struct {
	DBO DBOperator
	In  *events.ResetPasswordRequested
}

func (rp ResetPassword) Reset() *events.PasswordReseted {

	// make out event
	var out events.PasswordReseted
	out.EventStatus = &events.Status{
		HttpCode:  http.StatusOK,
		Timestamp: ptypes.TimestampNow(),
	}
	out.Uid = rp.In.Uid
	// make out event

	// validate and verify reset jwt
	if ok, err := jwtkit.ValidateExpired(jwtkit.JWTString(rp.In.Token)); !ok {
		log.Println(err)
		out.EventStatus.Errors = []string{err.Error()}
		out.EventStatus.HttpCode = http.StatusUnauthorized
		out.EventStatus.Timestamp = ptypes.TimestampNow()

		return &out
	}

	ecdsa := jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY"),
	}
	if ok, err := jwtkit.VerifyJWTString(&ecdsa, jwtkit.JWTString(rp.In.Token)); !ok {
		log.Println(err)
		out.EventStatus.Errors = []string{err.Error()}
		out.EventStatus.HttpCode = http.StatusUnauthorized
		out.EventStatus.Timestamp = ptypes.TimestampNow()

		return &out
	}
	// validate and verify reset jwt

	// get user uuid from token
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(rp.In.Token))
	errorkit.ErrorHandled(err)

	userUUID := jwt.Payload.Claims["user_uuid"].(string)
	// get user uuid from token

	// match given token with the one in db
	row, err := rp.DBO.QueryRow("SELECT rp.id,u.email FROM reset_passwords rp JOIN users u ON rp.user_uuid=u.uuid WHERE rp.user_uuid=? AND rp.token=?", userUUID, rp.In.Token)
	errorkit.ErrorHandled(err)

	var resetPasswordID uint64
	if err = row.Scan(&resetPasswordID, &out.Email); err == sql.ErrNoRows {
		log.Println(err)
		out.EventStatus.Errors = []string{"user with this email hasn't requested for password reset"}
		out.EventStatus.HttpCode = http.StatusNotFound
		out.EventStatus.Timestamp = ptypes.TimestampNow()

		return &out
	}
	// match given token with the one in db

	// update password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rp.In.NewPassword), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	log.Printf("hashed new password = %s", hashedPassword)
	err = rp.DBO.Command("UPDATE users SET password=? WHERE uuid=?;", hashedPassword, userUUID)
	errorkit.ErrorHandled(err)
	// update password

	return &out
}
