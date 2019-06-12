package grpc

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/user"
	"github.com/ilhammhdd/kudaki-externals/mysql"

	"github.com/golang/protobuf/ptypes"

	"github.com/ilhammhdd/go-toolkit/jwtkit"

	"github.com/ilhammhdd/kudaki-entities/events"
	kudakigrpc "github.com/ilhammhdd/kudaki-externals/grpc"
)

type User struct{}

func (u User) ResetPassword(context.Context, *events.ResetPasswordRequested) (*events.ResetPasswordEmailSent, error) {

	return nil, nil
}

func (u User) Signup(context.Context, *events.SignupRequested) (*events.Signedup, error) {
	return nil, nil
}

func (u User) VerifyUser(context.Context, *events.VerifyUserRequested) (*events.UserVerified, error) {
	return nil, nil
}

func (u User) Login(context.Context, *events.LoginRequested) (*events.Loggedin, error) {
	return nil, nil
}

func (u User) UserAuthentication(ctx context.Context, uar *kudakigrpc.UserAuthenticationRequested) (*kudakigrpc.UserAuthenticated, error) {
	ua := kudakigrpc.UserAuthenticated{
		Uid:         uar.Uid,
		EventStatus: new(events.Status)}

	userFromKudakiToken := u.getUserFromKudakiToken(uar.Jwt)
	if _, ok := u.userExists(userFromKudakiToken); !ok {
		ua.EventStatus.HttpCode = http.StatusNotFound
		ua.EventStatus.Timestamp = ptypes.TimestampNow()
		ua.EventStatus.Errors = []string{"user not found"}

		return &ua, nil
	}

	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}

	ok, err := jwtkit.VerifyJWTString(e, jwtkit.JWTString(uar.Jwt))

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

func (u User) getUserFromKudakiToken(kudakiToken string) *user.User {
	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(kudakiToken))
	errorkit.ErrorHandled(err)

	userClaim := jwt.Payload.Claims["user"].(map[string]interface{})
	usr := &user.User{
		AccountType: user.AccountType(user.AccountType_value[userClaim["account_type"].(string)]),
		Email:       userClaim["email"].(string),
		PhoneNumber: userClaim["phone_number"].(string),
		Role:        user.Role(user.Role_value[userClaim["role"].(string)]),
		Uuid:        userClaim["uuid"].(string),
	}

	return usr
}

func (u User) userExists(usr *user.User) (*user.User, bool) {
	dbo := mysql.NewDBOperation()
	row, err := dbo.QueryRow("SELECT uuid,email,password,token,role,phone_number,account_type FROM users WHERE uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var retrievedUser user.User
	var role string
	var accountType string
	if row.Scan(&retrievedUser.Uuid, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.Token, &role, &retrievedUser.PhoneNumber, &accountType) == sql.ErrNoRows {
		return nil, false
	}
	retrievedUser.Role = user.Role(user.Role_value[role])
	retrievedUser.AccountType = user.AccountType(user.AccountType_value[accountType])

	return &retrievedUser, true
}

func (u User) ChangePassword(ctx context.Context, rpp *events.ChangePasswordRequested) (*events.PasswordChanged, error) {

	return nil, nil
}

func (u User) UserAuthorization(ctx context.Context, uar *kudakigrpc.UserAuthorizationRequested) (*kudakigrpc.UserAuthorized, error) {

	jwt, err := jwtkit.GetJWT(jwtkit.JWTString(uar.Jwt))
	errorkit.ErrorHandled(err)

	dbo := mysql.NewDBOperation()
	row, err := dbo.QueryRow("SELECT id FROM users WHERE uuid=? AND role=?", jwt.Payload.Claims["user"].(map[string]interface{})["uuid"], user.Role_name[int32(uar.Role)])
	errorkit.ErrorHandled(err)

	var totalUserId int

	if scanErr := row.Scan(&totalUserId); scanErr == sql.ErrNoRows {
		grpcErr := "user's role not authorized"

		return &kudakigrpc.UserAuthorized{
			EventStatus: &events.Status{
				Errors:    []string{grpcErr},
				HttpCode:  http.StatusUnauthorized,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	} else {
		return &kudakigrpc.UserAuthorized{
			EventStatus: &events.Status{
				HttpCode:  http.StatusOK,
				Timestamp: ptypes.TimestampNow()},
			Uid: uar.Uid,
		}, nil
	}
}
