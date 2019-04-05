package grpc

import (
	"context"
	"net/http"
	"os"

	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-entities/events"
)

type User struct{}

func (u User) Signup(context.Context, *events.SignupRequested) (*events.UserVerificationEmailSent, error) {
	return nil, nil
}

func (u User) VerifyUser(context.Context, *events.VerifyUserRequested) (*events.Signedup, error) {
	return nil, nil
}

func (u User) Login(context.Context, *events.LoginRequested) (*events.Loggedin, error) {
	return nil, nil
}

func (u User) UserAuthentication(ctx context.Context, uar *events.UserAuthenticationRequested) (*events.UserAuthenticated, error) {

	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	ok, err := jwtkit.VerifyJWTString(e, jwtkit.JWTString(uar.Jwt))

	ua := events.UserAuthenticated{
		Uid:         uar.Uid,
		EventStatus: new(events.Status)}

	if !ok {
		ua.EventStatus.HttpCode = http.StatusUnauthorized
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"jwt not verified"}
	} else {
		ua.EventStatus.HttpCode = http.StatusOK
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"jwt verified"}
	}

	return &ua, err
}
