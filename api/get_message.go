package handler

import (
	"encoding/json"
	"jmessage/dbhelper"
	"math/rand"
	"net/http"
	"time"
)

func GetRandomChapter(w http.ResponseWriter, r *http.Request) {
	dbhelper.InitDB()

	// Zufälliges Kapitel auswählen
	var chapterID, bookID, chapterNumber int
	var bookName string
	query := `
		SELECT chapters.id, books.id, books.name, chapters.number
		FROM chapters
		INNER JOIN books ON chapters.book_id = books.id
		ORDER BY RANDOM() LIMIT 1;
	`
	err := dbhelper.DB.QueryRow(query).Scan(&chapterID, &bookID, &bookName, &chapterNumber)
	if err != nil {
		http.Error(w, "Failed to fetch random chapter", http.StatusInternalServerError)
		return
	}

	// Verse des Kapitels abrufen
	rows, err := dbhelper.DB.Query(`
		SELECT number, text
		FROM verses
		WHERE chapter_id = ? ORDER BY number;
	`, chapterID)
	if err != nil {
		http.Error(w, "Failed to fetch verses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	verses := []map[string]interface{}{}
	for rows.Next() {
		var number int
		var text string
		rows.Scan(&number, &text)
		verses = append(verses, map[string]interface{}{
			"number": number,
			"text":   text,
		})
	}

	// JSON-Antwort
	response := map[string]interface{}{
		"book":    bookName,
		"chapter": chapterNumber,
		"verses":  verses,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetRandomVerse(w http.ResponseWriter, r *http.Request) {
	dbhelper.InitDB()

	// Zufälligen Vers auswählen
	var verseNumber, chapterNumber int
	var text, bookName string
	query := `
		SELECT verses.number, chapters.number, books.name, verses.text
		FROM verses
		INNER JOIN chapters ON verses.chapter_id = chapters.id
		INNER JOIN books ON chapters.book_id = books.id
		ORDER BY RANDOM() LIMIT 1;
	`
	err := dbhelper.DB.QueryRow(query).Scan(&verseNumber, &chapterNumber, &bookName, &text)
	if err != nil {
		http.Error(w, "Failed to fetch random verse", http.StatusInternalServerError)
		return
	}

	// JSON-Antwort
	response := map[string]interface{}{
		"book":    bookName,
		"chapter": chapterNumber,
		"verse":   verseNumber,
		"text":    text,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
