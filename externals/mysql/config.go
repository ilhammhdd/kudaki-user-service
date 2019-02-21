package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ilhammhdd/go_error"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhammhdd/kudaki-user-service/entities"
)

func OpenDB(mysqlUser, mysqlPassword, mysqlPath, dbName string) {
	dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlPath, dbName)
	initDB, err := sql.Open("mysql", dbDataSource)
	go_error.HandleError(err)

	entities.DB = initDB
}
