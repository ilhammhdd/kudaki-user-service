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

func ResetPassword(dbo usecases.DBOperator, msg []byte) (key string, value []byte, err error) {
	var rpr events.ResetPasswordRequested
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
		passwordReseted := usecases.ResetPassword(&rpr, dbo)

		passwordResetedBytes, marshalErr := proto.Marshal(passwordReseted)
		return passwordReseted.Uid, passwordResetedBytes, marshalErr
	}

	return "", nil, unmarshalErr
}
