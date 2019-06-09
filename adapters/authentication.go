package adapters

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Signup struct{}

func (s *Signup) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.SignupRequested

	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		return &inEvent, true
	}

	return nil, false
}

func (s *Signup) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.Signedup)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type Login struct{}

func (l *Login) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.LoginRequested

	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		return &inEvent, true
	}

	return nil, false
}

func (l *Login) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.Loggedin)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type VerifyUser struct{}

func (vu *VerifyUser) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.VerifyUserRequested

	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		return &inEvent, true
	}

	return nil, false
}

func (vu *VerifyUser) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.UserVerified)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type ChangePassword struct{}

func (cp *ChangePassword) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ChangePasswordRequested
	if err := proto.Unmarshal(msg, &inEvent); err == nil {
		log.Printf("parsed ChangePasswordRequested : %v", inEvent)
		return &inEvent, true
	}
	return nil, false
}

func (cp *ChangePassword) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.PasswordChanged)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
