package eventdriven

import (
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"
	"github.com/ilhammhdd/kudaki-entities/user"
	"github.com/ilhammhdd/kudaki-externals/kafka"
	"github.com/ilhammhdd/kudaki-externals/mysql"
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
	s.insertUserAndProfile(out.Profile.User, out.Profile)
	s.indexUserAndProfile(out.Profile.User, out.Profile)
}

func (s *Signup) insertUserAndProfile(usr *user.User, profile *user.Profile) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES (?,?,?,?,?,?,?);",
		usr.Uuid, usr.Email, usr.Password, usr.Token, usr.Role.String(), usr.PhoneNumber, usr.AccountType.String())
	errorkit.ErrorHandled(err)

	_, err = dbo.Command("INSERT INTO profiles(uuid,user_uuid,full_name,photo,reputation) VALUES (?,?,?,?,?);",
		profile.Uuid, profile.User.Uuid, profile.FullName, profile.Photo, profile.Reputation)
	errorkit.ErrorHandled(err)
}

func (s *Signup) indexUserAndProfile(usr *user.User, profile *user.Profile) {
	userClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	userClient.CreateIndex(kudakiredisearch.User.Schema())

	userDoc := redisearch.NewDocument(kudakiredisearch.RedisearchText(usr.Uuid).Sanitize(), 1.0)
	userDoc.Set("user_uuid", usr.Uuid)
	userDoc.Set("user_email", usr.Email)
	userDoc.Set("user_password", usr.Password)
	userDoc.Set("user_token", usr.Token)
	userDoc.Set("user_role", usr.Role.String())
	userDoc.Set("user_phone_number", usr.PhoneNumber)
	userDoc.Set("user_account_type", usr.AccountType.String())

	err := userClient.IndexOptions(redisearch.DefaultIndexingOptions, userDoc)
	errorkit.ErrorHandled(err)

	profileClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Profile.Name())
	profileClient.CreateIndex(kudakiredisearch.Profile.Schema())

	profileDoc := redisearch.NewDocument(kudakiredisearch.RedisearchText(profile.Uuid).Sanitize(), 1.0)
	profileDoc.Set("profile_uuid", profile.Uuid)
	profileDoc.Set("profile_full_name", profile.FullName)
	profileDoc.Set("profile_photo", profile.Photo)
	profileDoc.Set("profile_reputation", profile.Reputation)

	err = profileClient.IndexOptions(redisearch.DefaultIndexingOptions, profileDoc)
	errorkit.ErrorHandled(err)
}
