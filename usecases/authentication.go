package usecases

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
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
		if errorkit.ErrorHandled(sendVerificationEmail(su)) {
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
			"user_uuid": lr.User.Uuid})

	errorkit.ErrorHandled(err)

	err = dbo.Command("UPDATE users SET token=? WHERE uuid=?", string(jwtString), usr.Uuid)
	errorkit.ErrorHandled(err)

	usr.Token = string(jwtString)
	produceLoggedin(esp, lr.Uid, nil, []string{"successfully logged in"}, http.StatusOK, &usr)
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

func sendVerificationEmail(su *events.SignupRequested) error {
	from := mail.Address{
		Name:    "Notification Kudaki.id",
		Address: os.Getenv("MAIL")}
	to := mail.Address{
		Name:    su.Profile.FullName,
		Address: su.Profile.User.Email}
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")

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

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = "User account verification"
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// servername := host + ":587"
	// auth := smtp.PlainAuth("", from.Address, password, host)

	// crtBytes, _ := ioutil.ReadFile("/certs/gateway.kudaki.id.crt")
	// keyBytes, _ := ioutil.ReadFile("/certs/gateway.kudaki.id.key")

	// certs, err := tls.X509KeyPair(crtBytes, keyBytes)
	// errorkit.ErrorHandled(err)

	// tlsConf := &tls.Config{
	// 	InsecureSkipVerify: false,
	// 	ServerName:         host,
	// 	Certificates:       []tls.Certificate{certs},
	// }

	// conn, err := tls.Dial("tcp", servername, tlsConf)
	// errorkit.ErrorHandled(err)

	// client, err := smtp.NewClient(conn, host)
	// errorkit.ErrorHandled(err)

	// err = client.Auth(auth)
	// errorkit.ErrorHandled(err)

	// err = client.Mail(from.Address)
	// errorkit.ErrorHandled(err)

	// err = client.Rcpt(to.Address)
	// errorkit.ErrorHandled(err)

	// mailWriter, err := client.Data()
	// errorkit.ErrorHandled(err)

	// _, err = mailWriter.Write([]byte(message))
	// errorkit.ErrorHandled(err)

	// err = mailWriter.Close()
	// errorkit.ErrorHandled(err)

	// client.Quit()

	auth := smtp.PlainAuth("", from.Address, password, host)
	err = smtp.SendMail(host+":25", auth, from.Address, []string{su.Profile.User.Email}, []byte(message))
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
}
