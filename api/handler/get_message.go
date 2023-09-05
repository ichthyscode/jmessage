package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"sync"
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

func GetMessage(w http.ResponseWriter, r *http.Request) {
	Mutex.Lock()
	defer Mutex.Unlock()

	var message string
	err := db.QueryRow("SELECT verse FROM messages WHERE time_to_show = CURRENT_DATE").Scan(&message)
	if err != nil {
		http.Error(w, "No verse found for today", 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
