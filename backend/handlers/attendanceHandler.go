package handlers

import (
	"Student-Attendance-System/backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func MarkAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

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

func ViewAttendanceReportHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

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
