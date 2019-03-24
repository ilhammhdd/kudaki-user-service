package usecases

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"

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
	row, err := dbOperator.QueryRow("SELECT count(id) FROM users WHERE email=?", su.Profile.User.Email)
	go_error.ErrorHandled(err)

	uves := events.UserVerificationEmailSent{
		Uuid:        su.Uuid,
		User:        su.Profile.User,
		EventStatus: &events.Status{},
	}

	var userID uint
	go_error.ErrorHandled(row.Scan(&userID))

	if userID > 0 {
		uves.EventStatus.Code = events.Code_BAD_COMMAND
		uves.EventStatus.Messages = []string{"user with the given email already exists"}
		uves.EventStatus.Source = entities.Services_USER
		uves.EventStatus.Timestamp = ptypes.TimestampNow()

		esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
		_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], &uves)
		go_error.ErrorHandled(err)

		return
	}

	if go_error.ErrorHandled(sendVerificationEmail(su)) {
		uves.EventStatus.Code = events.Code_INTERNAL_ERROR
		uves.EventStatus.Messages = []string{"error occured when sending verification email"}
		uves.EventStatus.Source = entities.Services_USER
		uves.EventStatus.Timestamp = ptypes.TimestampNow()

		esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
		_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], &uves)
		go_error.ErrorHandled(err)

		return
	}

	createUserAndProfile(su, dbOperator)

	uves.EventStatus.Code = events.Code_SUCCESS
	uves.EventStatus.Messages = []string{"successfully sent verfication email"}
	uves.EventStatus.Source = entities.Services_USER
	uves.EventStatus.Timestamp = ptypes.TimestampNow()

	esp.Set(entities.Topics_name[int32(entities.Topics_USER)], int32(entities.Partition_EVENT), sarama.OffsetNewest)
	_, _, err = esp.SyncProduce(entities.Partition_name[int32(entities.Partition_EVENT)], &uves)
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

	body := string(jwtString)

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
	password, err := bcrypt.GenerateFromPassword([]byte(su.Profile.User.Password), bcrypt.MinCost)
	go_error.ErrorHandled(err)

	dbo.Command(
		"INSERT INTO users(uuid,email,password,role,phone_number) VALUES(?,?,?,?,?)",
		su.Profile.User.Uuid,
		su.Profile.User.Email,
		password,
		user.Role_name[int32(su.Profile.User.Role)],
		su.Profile.User.PhoneNumber,
	)

	dbo.Command(
		"INSERT INTO unverified_users(user_uuid) VALUES(?)",
		su.Profile.User.Uuid,
	)
}
