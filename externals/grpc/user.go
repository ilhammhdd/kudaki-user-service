package grpc

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/user"
	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"

	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-entities/events"
)

type User struct{}

func (u User) Signup(context.Context, *events.SignupRequested) (*events.Signedup, error) {
	return nil, nil
}

func (u User) VerifyUser(context.Context, *events.VerifyUserRequested) (*events.UserVerified, error) {
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

func (u User) ResetPassword(ctx context.Context, rpp *events.ResetPasswordRequested) (*events.PasswordReseted, error) {

	return nil, nil
}

func (u User) UserAuthorization(ctx context.Context, uar *events.UserAuthorizationRequested) (*events.UserAuthorized, error) {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(uar.Jwt))
	errorkit.ErrorHandled(err)

	dbo := mysql.NewDBOperation()
	row, err := dbo.QueryRow("SELECT id FROM users WHERE uuid=? AND role=?", jwt.Payload.Claims["user"].(map[string]interface{})["uuid"], user.Role_name[int32(uar.Role)])
	errorkit.ErrorHandled(err)

	var totalUserId int

	if scanErr := row.Scan(&totalUserId); scanErr == sql.ErrNoRows {
		grpcErr := "user's role not authorized"

		return &events.UserAuthorized{
			EventStatus: &events.Status{
				Errors:    []string{grpcErr},
				HttpCode:  http.StatusUnauthorized,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	} else {
		return &events.UserAuthorized{
			EventStatus: &events.Status{
				HttpCode:  http.StatusOK,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	}
}
