package main

import (
	"Student-Attendance-System/backend/utils" // Import the utils package
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User struct represents the structure of the user we will add
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func main() {
	// Initialize the database
	utils.InitDB()

	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/addUser", addUserHandler)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}

// homeHandler handles the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Student Attendance System!")
}

// addUserHandler handles the adding of a new user
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Decode the JSON body into a User struct
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println("Failed to decode request body:", err)
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}

		// Call the function to add the user to the database
		err = utils.AddUser(user.Username, user.Password, user.Role)
		if err != nil {
			log.Println("Failed to add user:", err)
			http.Error(w, "Failed to add user", http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User added successfully")
	} else {
		// Handle non-POST requests
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
