package main

import (
	"database/sql"
	"log"

	"jmessage/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("db nil")
	}
	return db
}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS messages(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		time_to_show TEXT,
		verse TEXT
	);
	`
	_, err := db.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize SQLite database
	db := initDB("messages.db")
	createTable(db)
	defer db.Close()

	// Initialize Gin
	r := gin.Default()

	// API to get message
	r.GET("/message", func(c *gin.Context) { handler.GetMessageHandler(c, db) })

	// API to add message
	r.POST("/add-message", func(c *gin.Context) { handler.AddMessageHandler(c, db) })

	r.Run(":8080")
}
