package dbhelper

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./lutherbibel_1545.sqlite") // Pfad zur neuen DB
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
