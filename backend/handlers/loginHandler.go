package handlers

import (
	"Student-Attendance-System/backend/utils"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS request for preflight check
	if r.Method == http.MethodOptions {
		enableCORS(w, r)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle the actual POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body to get the login credentials
	var loginDetails struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the user exists and the password is correct
	user, err := utils.GetUserByUsername(loginDetails.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify the password
	if user.Password != loginDetails.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Login success - Send JSON response
	response := map[string]interface{}{
		"success": true,
		"role":    user.Role,
	}

	enableCORS(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
