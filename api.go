// api.go

package main

import (
	"net/http"

	"jmessage/api/handler"
)

func main() {
	handler.InitDB("your-connection-string-here")

	http.HandleFunc("/api/add_message", handler.AddMessage)
	http.HandleFunc("/api/get_message", handler.GetMessage)

	http.ListenAndServe(":8080", nil)
}
