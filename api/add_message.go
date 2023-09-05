package handler

import (
	"encoding/json"
	"net/http"
)

func AddMessage(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Date  string `json:"date"`
		Verse string `json:"verse"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	_, err := db.Exec("INSERT INTO messages(time_to_show, verse) VALUES($1, $2)", input.Date, input.Verse)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Verse added successfully"})
}
