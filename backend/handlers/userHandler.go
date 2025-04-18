package handlers

import (
	"Student-Attendance-System/backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body to get user details
	var userDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add the user to the database
	err = utils.AddUser(userDetails.Username, userDetails.Password, userDetails.Role)
	if err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User added successfully")
}
