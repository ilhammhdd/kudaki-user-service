package main

import (
	"os"

	"github.com/ilhammhdd/go_tool/go_error"

	"github.com/ilhammhdd/kudaki-user-service/externals/eventsourcing"

	"github.com/ilhammhdd/kudaki-user-service/externals/mysql"

	"github.com/ilhammhdd/go_tool/go_jwt"
	"github.com/ilhammhdd/go_tool/go_safe"
)

func init() {
	if len(os.Args) == 11 {
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
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	initJWT()
}

func initJWT() {
	e := &go_jwt.ECDSA{
		PrivateKeyPath: os.Getenv("VERIFICATION_PRIVATE_KEY"),
		PublicKeyPath:  os.Getenv("VERIFICATION_PUBLIC_KEY")}
	go_error.ErrorHandled(go_jwt.GeneratePublicPrivateToPEM(e))
}

func main() {
	wp := go_safe.NewWorkerPool()

	wp.Work <- eventsourcing.Signup

	wp.PoolWG.Wait()
}
