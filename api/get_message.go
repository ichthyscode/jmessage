package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var message string
	err := DB.QueryRow("SELECT verse FROM messages WHERE time_to_show = CURRENT_DATE").Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No verse found for today", 404)
		} else {
			http.Error(w, "Internal Server Error", 500)
		}
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
