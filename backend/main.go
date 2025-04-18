package main

import (
	"Student-Attendance-System/backend/handlers"
	"Student-Attendance-System/backend/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	utils.InitDB()

	// Set up HTTP routes
	http.HandleFunc("/login", handlers.LoginHandler)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
