package main

import (
	"jmessage/handler" // Replace with your module name if it's different

	"github.com/gin-gonic/gin"
)

func main() {
	handler.InitDB("messages.db") // Initialize the DB

	r := gin.Default()
	r.GET("/message", handler.GetMessage)
	r.POST("/add-message", handler.AddMessage)
	r.Run(":8080")
}
