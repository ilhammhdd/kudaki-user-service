package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ilhammhdd/kudaki-user-service/externals/eventsourcing"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/kudaki-entities/rpc"

	"google.golang.org/grpc"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"

	external_grpc "github.com/ilhammhdd/kudaki-user-service/externals/grpc"

	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"
)

func init() {
	if len(os.Args) == 14 {
		os.Setenv("KAFKA_BROKERS", os.Args[1])
		os.Setenv("DB_PATH", os.Args[2])
		os.Setenv("DB_USERNAME", os.Args[3])
		os.Setenv("DB_PASSWORD", os.Args[4])
		os.Setenv("DB_NAME", os.Args[5])
		os.Setenv("MAIL", os.Args[6])
		os.Setenv("MAIL_PASSWORD", os.Args[7])
		os.Setenv("MAIL_HOST", os.Args[8])
		os.Setenv("VERIFICATION_PRIVATE_KEY", os.Args[9])
		os.Setenv("VERIFICATION_PUBLIC_KEY", os.Args[10])
		os.Setenv("GATEWAY_HOST", os.Args[11])
		os.Setenv("GRPC_PORT", os.Args[12])
		os.Setenv("KAFKA_VERSION", os.Args[13])
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	initJWT()
	initAdmin()
}

func initJWT() {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(e))
}

func initAdmin() {
	if adminExists() {
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte("OlahragaOtak2K19!"), bcrypt.MinCost)
	errorkit.ErrorHandled(err)

	dbo := mysql.NewDBOperation()
	err = dbo.Command(
		"INSERT INTO users(uuid,email,password,token,role,phone_number,account_type) VALUES(?,?,?,?,?,?,?)",
		uuid.New().String(),
		"service@kudaki.id",
		password,
		"",
		user.Role_name[int32(user.Role_ADMIN)],
		"",
		user.AccountType_name[int32(user.AccountType_NATIVE)])
	errorkit.ErrorHandled(err)
}

func adminExists() bool {
	dbo := mysql.NewDBOperation()
	row, err := dbo.QueryRow(
		"SELECT count(id) FROM users WHERE role=?",
		user.Role_name[int32(user.Role_ADMIN)])
	errorkit.ErrorHandled(err)

	var totalIds int
	row.Scan(&totalIds)

	return totalIds == 1
}

func grpcListener() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	errorkit.ErrorHandled(err)

	grpcServer := grpc.NewServer()
	rpc.RegisterUserServer(grpcServer, external_grpc.User{})
	errorkit.ErrorHandled(grpcServer.Serve(lis))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- eventsourcing.Signup
	wp.Work <- eventsourcing.VerifyUser
	wp.Work <- eventsourcing.Login
	wp.Work <- eventsourcing.ResetPassword
	wp.Work <- grpcListener

	wp.PoolWG.Wait()
}
