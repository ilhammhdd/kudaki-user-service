package eventdriven

import (
	"net/http"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/user"
	"github.com/ilhammhdd/kudaki-externals/kafka"
	"github.com/ilhammhdd/kudaki-externals/mysql"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-user-service/adapters"
	"github.com/ilhammhdd/kudaki-user-service/usecases"
)

type Signup struct{}

func (s *Signup) Work() interface{} {
	usecase := &usecases.Signup{DBO: mysql.NewDBOperation(), Producer: kafka.NewProduction()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: s,
		eventDrivenAdapter:  new(adapters.Signup),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserTopic_SIGN_UP_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_SIGN_UP_REQUESTED.String()},
		outTopic:            events.UserTopic_SIGNED_UP.String()}
	ede.handle()
	return nil
}

func (s *Signup) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.Signedup)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	s.insertUser(out.Profile.User)
	s.insertProfile(out.Profile)
	s.indexUser(out.Profile.User)
	s.indexProfile(out.Profile.User.Uuid, out.Profile)
}

func (s *Signup) insertUser(usr *user.User) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES (?,?,?,?,?,?,?);",
		usr.Uuid, usr.Email, usr.Password, usr.Token, usr.Role.String(), usr.PhoneNumber, usr.AccountType.String())
	errorkit.ErrorHandled(err)

	_, err = dbo.Command("INSERT INTO unverified_users(user_uuid) VALUES (?);", usr.Uuid)
	errorkit.ErrorHandled(err)
}

func (s *Signup) insertProfile(profile *user.Profile) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO profiles(uuid,user_uuid,full_name,photo,reputation) VALUES (?,?,?,?,?);",
		profile.Uuid, profile.User.Uuid, profile.FullName, profile.Photo, profile.Reputation)
	errorkit.ErrorHandled(err)
}

func (s *Signup) indexUser(usr *user.User) {
	userClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	userClient.CreateIndex(kudakiredisearch.User.Schema())

	sanitizedUserUUID := kudakiredisearch.RedisearchText(usr.Uuid).Sanitize()
	userDoc := redisearch.NewDocument(sanitizedUserUUID, 1.0)
	userDoc.Set("user_uuid", sanitizedUserUUID)
	userDoc.Set("user_email", usr.Email)
	userDoc.Set("user_password", usr.Password)
	userDoc.Set("user_token", usr.Token)
	userDoc.Set("user_role", usr.Role.String())
	userDoc.Set("user_phone_number", usr.PhoneNumber)
	userDoc.Set("user_account_type", usr.AccountType.String())

	err := userClient.IndexOptions(redisearch.DefaultIndexingOptions, userDoc)
	errorkit.ErrorHandled(err)
}

func (s *Signup) indexProfile(userUUID string, profile *user.Profile) {
	profileClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Profile.Name())
	profileClient.CreateIndex(kudakiredisearch.Profile.Schema())

	sanitizedProfileUUID := kudakiredisearch.RedisearchText(profile.Uuid).Sanitize()
	profileDoc := redisearch.NewDocument(sanitizedProfileUUID, 1.0)
	profileDoc.Set("profile_uuid", sanitizedProfileUUID)
	profileDoc.Set("user_uuid", kudakiredisearch.RedisearchText(userUUID).Sanitize())
	profileDoc.Set("profile_full_name", profile.FullName)
	profileDoc.Set("profile_photo", profile.Photo)
	profileDoc.Set("profile_reputation", profile.Reputation)

	err := profileClient.IndexOptions(redisearch.DefaultIndexingOptions, profileDoc)
	errorkit.ErrorHandled(err)
}

type Login struct{}

func (l *Login) Work() interface{} {
	usecase := &usecases.Login{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: l,
		eventDrivenAdapter:  new(adapters.Login),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserTopic_LOGIN_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_LOGIN_REQUESTED.String()},
		outTopic:            events.UserTopic_LOGGED_IN.String()}
	ede.handle()
	return nil
}

func (l *Login) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.Loggedin)
	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}
	l.updateUserAuthToken(out.User)
	l.reindexUser(out.User)
}

func (l *Login) updateUserAuthToken(usr *user.User) {
	dbo := mysql.NewDBOperation()

	_, err := dbo.Command("UPDATE users SET token=? WHERE uuid=?;", usr.Token, usr.Uuid)
	errorkit.ErrorHandled(err)
}

func (l *Login) reindexUser(usr *user.User) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	client.CreateIndex(kudakiredisearch.User.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(usr.Uuid).Sanitize(), 1.0)
	doc.Set("user_token", usr.Token)
	client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
}

type VerifyUser struct{}

func (vu *VerifyUser) Work() interface{} {
	usecase := &usecases.VerifyUser{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: vu,
		eventDrivenAdapter:  new(adapters.VerifyUser),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserTopic_VERIFY_USER_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_VERIFY_USER_REQUESTED.String()},
		outTopic:            events.UserTopic_USER_VERIFIED.String()}
	ede.handle()
	return nil
}

func (vu *VerifyUser) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.UserVerified)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("DELETE FROM unverified_users WHERE user_uuid = ?;", out.User.Uuid)
	errorkit.ErrorHandled(err)
}

type ChangePassword struct{}

func (cp *ChangePassword) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.PasswordChanged)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	cp.updateUsersPassword(out.User)
	cp.reIndexUser(out.User)
}

func (cp *ChangePassword) updateUsersPassword(usr *user.User) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE users SET password = ? WHERE uuid = ?;", usr.Password, usr.Uuid)
	errorkit.ErrorHandled(err)
}

func (cp *ChangePassword) reIndexUser(usr *user.User) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	client.CreateIndex(kudakiredisearch.User.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(usr.Uuid).Sanitize(), 1.0)
	doc.Set("user_password", usr.Password)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

func (cp *ChangePassword) Work() interface{} {
	usecase := &usecases.ChangePassword{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: cp,
		eventDrivenAdapter:  new(adapters.ChangePassword),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserTopic_CHANGE_PASSWORD_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_CHANGE_PASSWORD_REQUESTED.String()},
		outTopic:            events.UserTopic_PASSWORD_CHANGED.String()}

	ede.handle()
	return nil
}

type ResetPasswordSendEmail struct{}

func (rpse *ResetPasswordSendEmail) Work() interface{} {
	usecase := usecases.ResetPasswordSendEmail{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rpse,
		eventDrivenAdapter:  new(adapters.ResetPasswordSendEmail),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED.String()},
		outTopic:            events.UserTopic_RESET_PASSWORD_EMAIL_SENT.String()}

	ede.handle()
	return nil
}

func (rpse *ResetPasswordSendEmail) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.ResetPasswordEmailSent)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO reset_passwords(user_uuid,token) VALUES(?,?) ON DUPLICATE KEY UPDATE token = ?;", out.User.Uuid, out.RestToken, out.User.Uuid)
	errorkit.ErrorHandled(err)
}

type ResetPassword struct{}

func (rp *ResetPassword) Work() interface{} {
	usecase := &usecases.ResetPassword{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rp,
		eventDrivenAdapter:  new(adapters.ResetPassword),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserTopic_RESET_PASSWORD_REQUESTED.String(),
		inTopics:            []string{events.UserTopic_RESET_PASSWORD_REQUESTED.String()},
		outTopic:            events.UserTopic_PASSWORD_RESETED.String()}

	ede.handle()
	return nil
}

func (rp *ResetPassword) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.PasswordReseted)
	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	rp.updateUser(out)
	rp.reIndexUser(out)
}

func (rp *ResetPassword) updateUser(out *events.PasswordReseted) {
	dbo := mysql.NewDBOperation()

	_, err := dbo.Command("UPDATE users SET password=? WHERE uuid = ?;", out.User.Password, out.User.Uuid)
	errorkit.ErrorHandled(err)
}

func (rp *ResetPassword) reIndexUser(out *events.PasswordReseted) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	client.CreateIndex(kudakiredisearch.User.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(out.User.Uuid).Sanitize(), 1.0)
	doc.Set("user_password", out.User.Password)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
	errorkit.ErrorHandled(err)
}
