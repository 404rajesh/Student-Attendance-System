package main

import (
	"Student-Attendance-System/backend/utils" // Import the utils package
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Enable CORS for all routes
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")                   // Allow methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow headers
}

// Handler for login with CORS handling
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS request for preflight check
	if r.Method == http.MethodOptions {
		enableCORS(w, r)             // Enable CORS for preflight
		w.WriteHeader(http.StatusOK) // Respond with status OK for OPTIONS request
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
		"role":    user.Role, // Send the role of the logged-in user
	}

	enableCORS(w, r) // Enable CORS for the actual login response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Initialize server and routes
func main() {
	// Initialize the database connection
	utils.InitDB()

	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/addUser", addUserHandler)
	http.HandleFunc("/login", loginHandler) // Handle login (includes CORS)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}

// Handler for the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r) // Enable CORS for this request
	fmt.Fprintf(w, "Welcome to the Student Attendance System!")
}

// Handler for adding a new user (sign-up)
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r) // Enable CORS for this request

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

// Handler for marking attendance
func markAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r) // Enable CORS for this request

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body to get the attendance details
	var attendanceDetails struct {
		UserID string `json:"user_id"`
		Status string `json:"status"`
		Date   string `json:"date"`
	}

	err := json.NewDecoder(r.Body).Decode(&attendanceDetails)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mark the attendance
	err = utils.MarkAttendance(attendanceDetails.UserID, attendanceDetails.Status, attendanceDetails.Date)
	if err != nil {
		http.Error(w, "Failed to mark attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Attendance marked successfully")
}

// Handler for viewing attendance report
func viewAttendanceReportHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r) // Enable CORS for this request

	// You could pass query params like student_id and date
	studentID := r.URL.Query().Get("student_id")
	date := r.URL.Query().Get("date")

	// Fetch the attendance data from the database
	attendance, err := utils.GetAttendanceReport(studentID, date)
	if err != nil {
		http.Error(w, "Failed to fetch attendance data", http.StatusInternalServerError)
		return
	}

	// Respond with the attendance data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendance)
}
