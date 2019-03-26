package adapters

import (
	"github.com/ilhammhdd/kudaki-entities/commands"
	"github.com/ilhammhdd/kudaki-user-service/usecases"

	"github.com/golang/protobuf/proto"

	"github.com/ilhammhdd/go_tool/go_error"
)

func Signup(dbOperator usecases.DBOperator, esp usecases.EventSourceProducer, msg []byte) {

	var signUp commands.Signup

	err := proto.Unmarshal(msg, &signUp)
	if !go_error.ErrorHandled(err) {
		usecases.Signup(&signUp, dbOperator, esp)
	}
}
