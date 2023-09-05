package dbhelper

import (
	"database/sql"
	"os"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("POSTGRES_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
