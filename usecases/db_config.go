package usecases

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	SourceName string `json:"source_name"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`
}

func (c *Config) OpenDB() (*sql.DB, error) {
	dbDataSource := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", c.User, c.Password, c.SourceName, c.Database)
	db, err := sql.Open("mysql", dbDataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
