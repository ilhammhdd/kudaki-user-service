package main

import (
	"os"

	"github.com/ilhammhdd/kudaki-user-service/adapters"

	"github.com/ilhammhdd/go_safe"

	"github.com/ilhammhdd/go_error"

	"github.com/ilhammhdd/go_db"

	"github.com/ilhammhdd/kudaki-user-service/entities"
)

func init() {
	if len(os.Args) == 6 {
		os.Setenv("KAFKA_SERVER", os.Args[1])
		os.Setenv("DB_PATH", os.Args[2])
		os.Setenv("DB_USERNAME", os.Args[3])
		os.Setenv("DB_PASSWORD", os.Args[4])
		os.Setenv("DB_NAME", os.Args[5])
	}

	entities.DBConfig = go_db.Config{
		SourceName: os.Getenv("DB_PATH"),
		User:       os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		Database:   os.Getenv("DB_NAME")}

	db, err := entities.DBConfig.OpenDB()
	entities.DB = db
	go_error.HandleError(err)

}

func main() {
	wp := go_safe.NewWorkerPool()
	wp.Job <- adapters.SignUp{}
	wp.PoolWG.Wait()
}
