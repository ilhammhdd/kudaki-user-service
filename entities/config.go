package entities

import (
	"database/sql"

	"github.com/ilhammhdd/go_db"
)

const (
	Address = "ADDRESS"
)

var DB *sql.DB

var DBConfig go_db.Config
