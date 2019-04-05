package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ilhammhdd/kudaki-entities/rpc"

	"google.golang.org/grpc"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/jwtkit"
	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/kudaki-user-service/externals/eventsourcing"
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
		os.Setenv("GRPC_ADDRESS", os.Args[12])
		os.Setenv("GRPC_PORT", os.Args[13])
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	initJWT()
}

func initJWT() {
	e := &jwtkit.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}
	errorkit.ErrorHandled(jwtkit.GeneratePublicPrivateToPEM(e))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Work <- eventsourcing.Signup
	wp.Work <- eventsourcing.VerifyUser
	wp.Work <- eventsourcing.Login
	wp.Work <- grpcListener

	wp.PoolWG.Wait()
}

func grpcListener() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GRPC_ADDRESS"), os.Getenv("GRPC_PORT")))
	errorkit.ErrorHandled(err)

	grpcServer := grpc.NewServer()
	rpc.RegisterUserServer(grpcServer, external_grpc.User{})
	errorkit.ErrorHandled(grpcServer.Serve(lis))
}
