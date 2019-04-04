package adapters

import (
	"log"

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
		log.Println("it's a LoginRequested event")
		usecases.Login(&loginRequested, dbOperator, esp)
	}
}
