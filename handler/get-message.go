package main

import (
	"database/sql"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var mutex = &sync.Mutex{}

func GetMessage(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	defer db.Close()

	var message string
	err = db.QueryRow("SELECT verse FROM messages WHERE time_to_show = CURRENT_DATE").Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "No verse found for today"})
		} else {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(200, gin.H{"message": message})
}
