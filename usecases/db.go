package usecases

import (
	"database/sql"
)

type DBOperator interface {
	Command(string, ...interface{}) error
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) (*sql.Row, error)
	QueryRowsToMap(string, ...interface{}) (*[]map[string]interface{}, error)
}

// type ProtoSql interface {
// 	ScanRow(*sql.Row) error
// 	ScanRows(*sql.Rows) error
// }
