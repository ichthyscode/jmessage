package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetMessageHandler(c *gin.Context, db *sql.DB) {
	mutex.Lock()
	defer mutex.Unlock()
	var message string
	err := db.QueryRow("SELECT verse FROM messages WHERE time_to_show = CURRENT_DATE").Scan(&message)
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
