package main

import (
	"Student-Attendance-System/backend/config"
	"Student-Attendance-System/backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	config.ConnectDB()

	// Initialize the router
	r := mux.NewRouter()

	// Set up routes
	routes.InitializeAuthRoutes(r)
	routes.InitializeAttendanceRoutes(r)

	// Start the server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
