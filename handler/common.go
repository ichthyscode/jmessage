package handler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Importing SQLite3 driver
)

var DB *sql.DB

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	DB = db
	return DB
}
