package handler

import (
	"database/sql"
	"sync"

	"github.com/gin-gonic/gin"
)

var Mutex = &sync.Mutex{}

func GetMessage(c *gin.Context) {
	Mutex.Lock()
	var message string
	err := DB.QueryRow("SELECT verse FROM messages WHERE time_to_show = CURRENT_DATE").Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "No verse found for today"})
		} else {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
	} else {
		c.JSON(200, gin.H{"message": message})
	}
	Mutex.Unlock()
}
