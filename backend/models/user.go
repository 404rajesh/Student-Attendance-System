package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system (student or teacher)
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null"` // 'student' or 'teacher'
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register function registers a new user
func (user *User) Register(db *gorm.DB) error {
	result := db.Create(&user)
	return result.Error
}

// FindUserByUsername finds a user by username
func FindUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
