package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ilhammhdd/go_tool/go_error"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB(sourceName, user, password, database string) {
	dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, sourceName, database)
	initDB, err := sql.Open("mysql", dbDataSource)
	go_error.ErrorHandled(err)

	DB = initDB
}
