package adapters

import (
	"log"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"gopkg.in/Shopify/sarama.v1"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-user-service/usecases"

	"github.com/golang/protobuf/proto"
)

func Signup(dbOperator usecases.DBOperator, esp usecases.EventDrivenProducer, msg []byte) {

	signUp := &events.SignupRequested{}

	err := proto.Unmarshal(msg, signUp)
	if err == nil {
		usecases.Signup(signUp, dbOperator, esp)
	}
}

func VerifyUser(dbOperator usecases.DBOperator, msg []byte) (key string, value []byte, err error) {

	var verifyUser events.VerifyUserRequested

	unmarshalErr := proto.Unmarshal(msg, &verifyUser)
	if unmarshalErr == nil {
		signedUp := usecases.VerifyUser(&verifyUser, dbOperator)
		signedUpBytes, marshalErr := proto.Marshal(signedUp)

		return signedUp.Uid, signedUpBytes, marshalErr
	}

	return "", nil, unmarshalErr
}

func Login(dbOperator usecases.DBOperator, msg *sarama.ConsumerMessage) (key string, value []byte, err error) {

	var loginRequested events.LoginRequested
	unmarshallErr := proto.Unmarshal(msg.Value, &loginRequested)
	if unmarshallErr == nil {
		log.Printf("consumed LoginRequested : key = %s, offset = %v, partition = %v", string(msg.Key), msg.Offset, msg.Partition)
		loggedIn := usecases.Login(&loginRequested, dbOperator)
		loggedInBytes, marshalErr := proto.Marshal(loggedIn)

		return loggedIn.Uid, loggedInBytes, marshalErr
	}

	return "", nil, unmarshallErr
}

func ChangePassword(dbo usecases.DBOperator, msg []byte) (key string, value []byte, err error) {
	var rpr events.ChangePasswordRequested
	var fullName string
	var email string

	unmarshalErr := proto.Unmarshal(msg, &rpr)
	if unmarshalErr == nil {

		gotJWT, _ := jwtkit.GetJWT(jwtkit.JWTString(rpr.Profile.User.Token))
		userClaims := gotJWT.Payload.Claims["user"].(map[string]interface{})

		row, err := dbo.QueryRow("SELECT full_name FROM profiles WHERE user_uuid=?", userClaims["uuid"])
		errorkit.ErrorHandled(err)

		row.Scan(&fullName)

		row, err = dbo.QueryRow("SELECT email FROM users WHERE uuid=?", userClaims["uuid"])
		errorkit.ErrorHandled(err)

		row.Scan(&email)

		rpr.Profile.FullName = fullName
		rpr.Profile.User.Email = email
		passwordChanged := usecases.ChangePassword(&rpr, dbo)

		passwordChangedBytes, marshalErr := proto.Marshal(passwordChanged)
		return passwordChanged.Uid, passwordChangedBytes, marshalErr
	}

	return "", nil, unmarshalErr
}

type SendResetPasswordEmail struct {
	Partition int32
	Offset    int64
	Key       string
	Message   *[]byte
}

func (pr SendResetPasswordEmail) SendEmail(dbo usecases.DBOperator) (key string, msg []byte, err error) {

	var rpr events.SendResetPasswordEmailRequested

	unmarshalErr := proto.Unmarshal(*pr.Message, &rpr)
	if unmarshalErr == nil {

		prUsecase := usecases.SendResetPasswordEmail{
			DBO: dbo,
			In:  &rpr,
		}
		rpes := prUsecase.SendEmail()
		rpesByte, err := proto.Marshal(rpes)
		errorkit.ErrorHandled(err)

		return rpes.Uid, rpesByte, nil
	}

	return "", nil, unmarshalErr
}

type ResetPassword struct {
	Partition int32
	Offset    int64
	Key       string
	Message   *[]byte
}

func (rp ResetPassword) Reset(dbo usecases.DBOperator) (string, []byte, error) {

	in, err := rp.parseKafkaMessageToEvent()
	if err == nil {
		rpUsecase := usecases.ResetPassword{
			DBO: dbo,
			In:  in,
		}
		key, val := rp.parseToKafkaMessage(rpUsecase.Reset())
		return key, val, nil
	}
	return "", nil, err
}

func (rp ResetPassword) parseKafkaMessageToEvent() (*events.ResetPasswordRequested, error) {

	var in events.ResetPasswordRequested

	err := proto.Unmarshal(*rp.Message, &in)
	if err == nil {
		return &in, nil
	}

	return nil, err
}

func (rp ResetPassword) parseToKafkaMessage(in *events.PasswordReseted) (string, []byte) {

	out, marshalErr := proto.Marshal(in)
	errorkit.ErrorHandled(marshalErr)

	return in.Uid, out
}
