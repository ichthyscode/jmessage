package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	fmt.Println("Successfully connected to DB")
}

func init() {
	// Read the database connection string from an environment variable
	connectionString := os.Getenv("POSTGRES_URL")

	// Initialize the database
	InitDB(connectionString)
}
