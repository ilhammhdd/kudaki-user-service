package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-user-service/externals/eventdriven"

	"github.com/RediSearch/redisearch-go/redisearch"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/user"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"

	kudakigrpc "github.com/ilhammhdd/kudaki-externals/grpc"

	"google.golang.org/grpc"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"

	external_grpc "github.com/ilhammhdd/kudaki-user-service/externals/grpc"

	"github.com/ilhammhdd/kudaki-externals/mysql"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(string(f[1]), f[2])
		}
	}

	// debugging
	log.Printf("MAIL = %s", os.Getenv("MAIL"))
	log.Printf("MAIL_PASSWORD = %s", os.Getenv("MAIL_PASSWORD"))
	log.Printf("MAIL_HOST = %s", os.Getenv("MAIL_HOST"))
	log.Printf("MAIL_PORT = %s", os.Getenv("MAIL_PORT"))
	// debugging

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	initJWT()
	initAdmin()
}

func initJWT() {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(e))

	ecdsa := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("RESET_PASSWORD_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("RESET_PASSWORD_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(ecdsa))
}

func initAdmin() {
	if adminExists() {
		log.Println("admin already exists")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte("OlahragaOtak2K19!"), bcrypt.MinCost)
	errorkit.ErrorHandled(err)
	uuid := uuid.New().String()

	dbo := mysql.NewDBOperation()
	result, err := dbo.Command(
		"INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES(?,?,?,?,?,?,?)",
		uuid, "kudaki.service@gmail.com", password, "", user.Role_name[int32(user.Role_ADMIN)], "", user.AccountType_name[int32(user.AccountType_NATIVE)])
	errorkit.ErrorHandled(err)
	userlastInsertedID, err := result.LastInsertId()

	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
	client.CreateIndex(kudakiredisearch.User.Schema())
	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(uuid).Sanitize(), 1.0)
	doc.Set("user_id", userlastInsertedID)
	doc.Set("user_uuid", kudakiredisearch.RedisearchText(uuid).Sanitize())
	doc.Set("user_email", kudakiredisearch.RedisearchText("kudaki.service@gmail.com").Sanitize())
	doc.Set("user_password", password)
	doc.Set("user_token", "")
	doc.Set("user_role", user.Role_ADMIN.String())
	doc.Set("user_phone_number", "")
	doc.Set("user_account_type", user.AccountType_NATIVE.String())
	client.IndexOptions(redisearch.DefaultIndexingOptions, doc)
}

func adminExists() bool {
	dbo := mysql.NewDBOperation()
	row, err := dbo.QueryRow(
		"SELECT count(id) FROM users WHERE role=?",
		user.Role_ADMIN.String())
	errorkit.ErrorHandled(err)

	var totalIds int
	row.Scan(&totalIds)

	return totalIds == 1
}

func grpcListener() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	errorkit.ErrorHandled(err)

	grpcServer := grpc.NewServer()
	kudakigrpc.RegisterUserServer(grpcServer, external_grpc.User{})
	errorkit.ErrorHandled(grpcServer.Serve(lis))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(eventdriven.Signup)
	wp.Worker <- new(eventdriven.Login)
	wp.Worker <- new(eventdriven.VerifyUser)
	wp.Worker <- new(eventdriven.ChangePassword)
	wp.Worker <- new(eventdriven.ResetPasswordSendEmail)
	wp.Worker <- new(eventdriven.ResetPassword)
	wp.Work <- grpcListener

	wp.PoolWG.Wait()
}
