package models

import (
	"time"

	"gorm.io/gorm"
)

// Attendance represents attendance entry for a user
type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Date      time.Time `json:"date"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Status    string    `json:"status"`     // 'Present' or 'Absent'
	ClassCode string    `json:"class_code"` // Optional, for manual code entry
}

// 👇 Tell GORM to use "attendance" table
func (Attendance) TableName() string {
	return "attendance"
}

// MarkAttendance function marks attendance for a user
func (a *Attendance) MarkAttendance(db *gorm.DB) error {
	result := db.Create(a)
	return result.Error
}

// FindAttendanceByUser finds all attendance records for a user
func FindAttendanceByUser(db *gorm.DB, userID uint) ([]Attendance, error) {
	var attendance []Attendance
	err := db.Where("user_id = ?", userID).Find(&attendance).Error
	if err != nil {
		return nil, err
	}
	return attendance, nil
}
