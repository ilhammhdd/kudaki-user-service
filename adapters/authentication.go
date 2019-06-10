package adapters

import (
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

type ResetPasswordSendEmail struct{}

func (rpse *ResetPasswordSendEmail) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.SendResetPasswordEmailRequested
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}

	return nil, false
}

func (rpse *ResetPasswordSendEmail) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ResetPasswordEmailSent)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type ResetPassword struct{}

func (rp *ResetPassword) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.ResetPasswordRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rp *ResetPassword) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.PasswordReseted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
