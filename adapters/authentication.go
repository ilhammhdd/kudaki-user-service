package adapters

import (
	"log"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-user-service/usecases"

	"github.com/golang/protobuf/proto"
)

func Signup(dbOperator usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {

	signUp := &events.SignupRequested{}

	err := proto.Unmarshal(msg, signUp)
	if err == nil {
		log.Println("It's a SignupRequested event")
		usecases.Signup(signUp, dbOperator, esp)
	}
}

func VerifyUser(dbOperator usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {

	verifyUser := &events.VerifyUserRequested{}

	err := proto.Unmarshal(msg, verifyUser)
	if err == nil {
		log.Println("It's a VerifyUserRequested event")
		usecases.VerifyUser(verifyUser, dbOperator, esp)
	}
}

func Login(dbOperator usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {

	var loginRequested events.LoginRequested

	err := proto.Unmarshal(msg, &loginRequested)
	if err == nil {
		log.Println("it's a LoginRequested event", loginRequested)
		usecases.Login(&loginRequested, dbOperator, esp)
	}
}

func ResetPassword(dbo usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {
	var rpr events.ResetPasswordRequested
	var fullName string
	var email string

	if err := proto.Unmarshal(msg, &rpr); err == nil {
		log.Println("it's a ResetPasswordRequested event", rpr)

		row, err := dbo.QueryRow("SELECT full_name FROM profiles WHERE user_uuid=?", rpr.Profile.User.Uuid)
		errorkit.ErrorHandled(err)

		row.Scan(&fullName)

		row, err = dbo.QueryRow("SELECT email FROM users WHERE uuid=?", rpr.Profile.User.Uuid)
		errorkit.ErrorHandled(err)

		row.Scan(&email)

		rpr.Profile.FullName = fullName
		rpr.Profile.User.Email = email

		log.Println("profile full_name :", fullName)
		log.Println("user email :", email)
		usecases.ResetPassword(&rpr, dbo, esp)
	}
}
