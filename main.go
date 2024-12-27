package main

import (
	"jmessage/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	// Neue Endpunkte
	r.HandleFunc("/api/random-chapter", handler.GetRandomChapter).Methods("GET")
	r.HandleFunc("/api/random-verse", handler.GetRandomVerse).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
