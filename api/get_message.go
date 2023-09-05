package handler

import (
	"encoding/json"
	"net/http"
)

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
