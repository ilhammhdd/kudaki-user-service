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
