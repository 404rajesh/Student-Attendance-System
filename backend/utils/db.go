package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the database
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./attendance.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Create users table
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL CHECK (role IN ('student', 'teacher', 'admin'))
	);
	`

	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	fmt.Println("✅ Database initialized successfully.")
}

// AddUser adds a new user to the database
func AddUser(username, password, role string) error {
	// Prepare the SQL statement
	stmt, err := DB.Prepare("INSERT INTO users (username, password, role) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(username, password, role)
	if err != nil {
		log.Println("Error executing statement:", err)
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
