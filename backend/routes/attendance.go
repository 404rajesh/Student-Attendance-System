package routes

import (
	"Student-Attendance-System/backend/config"
	"Student-Attendance-System/backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// validateToken is a utility function to check JWT token validity and return the user ID
func validateToken(tokenString string) (uint, error) {
	// Remove Bearer prefix if exists
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, fmt.Errorf("invalid token")
}

// MarkAttendance marks the attendance for a student
func MarkAttendance(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}

	userID, err := validateToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	attendance.Date = time.Now()
	attendance.Time = time.Now()
	attendance.Status = "Present"
	attendance.UserID = userID

	if err := attendance.MarkAttendance(config.DB); err != nil {
		http.Error(w, "Error marking attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendance)
}

// ViewAttendance returns attendance records for a student
func ViewAttendance(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}

	userID, err := validateToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	requestedUserID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if uint(requestedUserID) != userID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	attendance, err := models.FindAttendanceByUser(config.DB, uint(userID))
	if err != nil {
		http.Error(w, "Error retrieving attendance", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(attendance)
}

// Initialize routes
func InitializeAttendanceRoutes(r *mux.Router) {
	r.HandleFunc("/attendance", MarkAttendance).Methods("POST")
	r.HandleFunc("/attendance/{user_id}", ViewAttendance).Methods("GET")
}
