package adapters

import (
	"log"

	"github.com/ilhammhdd/kudaki-user-service/entities/commands"
	"github.com/ilhammhdd/kudaki-user-service/usecases"

	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go_tool/go_error"
)

func Signup(dbOperator usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {

	var signUp commands.SignUp

	err := proto.Unmarshal(msg, &signUp)
	if !go_error.ErrorHandled(err) {
		log.Println("it's a SignUp command : ", signUp)
		usecases.Signup(&signUp, dbOperator, esp)
	}

	log.Println("it's not a SignUp command")
}
