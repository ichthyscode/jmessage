package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// API routes
	r.GET("/message", GetMessage)
	r.POST("/add-message", AddMessage)

	// Run server
	r.Run(":8080")
}
