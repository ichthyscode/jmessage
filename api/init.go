package handler

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	Mutex sync.Mutex
)

func init() {
	connStr := os.Getenv("POSTGRES_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
