package usecases

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go_tool/go_jwt"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/Shopify/sarama.v1"

	"github.com/golang/protobuf/ptypes"

	entities "github.com/ilhammhdd/kudaki-entities"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/go_tool/go_error"
	"github.com/ilhammhdd/kudaki-entities/commands"
)

func Signup(su *commands.Signup, dbOperator DBOperator, esp EventSourceProducer) {
	password, err := bcrypt.GenerateFromPassword([]byte(su.Profile.User.Password), bcrypt.MinCost)
	go_error.ErrorHandled(err)
	su.Profile.User.Password = string(password)

	var eventStatus events.Status
	uves := events.UserVerificationEmailSent{
		Uuid: su.Uuid,
		User: su.Profile.User,
	}

	row, err := dbOperator.QueryRow("SELECT count(id) FROM users WHERE email=?", su.Profile.User.Email)
	go_error.ErrorHandled(err)

	var userID uint
	go_error.ErrorHandled(row.Scan(&userID))

	if userID > 0 {
		eventStatus.HttpCode = http.StatusConflict
		eventStatus.Errors = []string{"user with the given email already exists"}
		eventStatus.Timestamp = ptypes.TimestampNow()
		uves.EventStatus = &eventStatus
		uvesBytes, err := proto.Marshal(&uves)
		go_error.ErrorHandled(err)

		esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
		_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], uvesBytes)
		go_error.ErrorHandled(err)

		return
	}

	if go_error.ErrorHandled(sendVerificationEmail(su)) {
		eventStatus.HttpCode = http.StatusBadRequest
		eventStatus.Errors = []string{"error occured when sending verification email"}
		eventStatus.Timestamp = ptypes.TimestampNow()
		uves.EventStatus = &eventStatus
		uvesBytes, err := proto.Marshal(&uves)
		go_error.ErrorHandled(err)

		esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
		_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], uvesBytes)
		go_error.ErrorHandled(err)

		return
	}

	createUserAndProfile(su, dbOperator)

	eventStatus.HttpCode = http.StatusOK
	eventStatus.Messages = []string{"successfully sent verfication email"}
	eventStatus.Timestamp = ptypes.TimestampNow()
	uves.EventStatus = &eventStatus
	uvesBytes, err := proto.Marshal(&uves)
	go_error.ErrorHandled(err)

	log.Println("UserVerificationEmailSent", uves.Uuid)

	esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
	_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], uvesBytes)
	go_error.ErrorHandled(err)
}

func sendVerificationEmail(su *commands.Signup) error {
	from := mail.Address{
		Name:    "Notification Kudaki.id",
		Address: os.Getenv("MAIL")}
	to := mail.Address{
		Name:    su.Profile.FullName,
		Address: su.Profile.User.Email}
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")

	e := &go_jwt.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	je := go_jwt.JWTExpiration(172800000)
	jwtString, err := je.GenerateSignedJWTString(e, "unverified Kudaki.id user", "Kudaki.id user service")
	go_error.ErrorHandled(err)

	body := "https://kudaki.id/user/verify/?verify_token=" + string(jwtString)

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

	auth := smtp.PlainAuth("", from.Address, password, host)
	err = smtp.SendMail(host+":587", auth, from.Address, []string{su.Profile.User.Email}, []byte(message))
	return err
}

func createUserAndProfile(su *commands.Signup, dbo DBOperator) {
	dbo.Command(
		"INSERT INTO users(uuid,email,password,role,phone_number) VALUES(?,?,?,?,?)",
		su.Profile.User.Uuid,
		su.Profile.User.Email,
		su.Profile.User.Password,
		user.Role_name[int32(su.Profile.User.Role)],
		su.Profile.User.PhoneNumber,
	)

	dbo.Command(
		"INSERT INTO unverified_users(user_uuid) VALUES(?)",
		su.Profile.User.Uuid,
	)
}
