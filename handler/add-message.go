package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func AddMessageHandler(c *gin.Context, db *sql.DB) {
	mutex.Lock()
	defer mutex.Unlock()
	var input struct {
		Date  string `json:"date"`
		Verse string `json:"verse"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	_, err := db.Exec("INSERT INTO messages(time_to_show, verse) VALUES(?, ?)", input.Date, input.Verse)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Verse added successfully"})
}
