package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the database
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

	// Create attendance table
	createAttendanceTable := `
	CREATE TABLE IF NOT EXISTS attendance (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		date TEXT NOT NULL,
		status TEXT CHECK (status IN ('present', 'absent')),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	_, err = DB.Exec(createAttendanceTable)
	if err != nil {
		log.Fatal("Failed to create attendance table:", err)
	}

	fmt.Println("✅ Database initialized successfully.")

	if err := InitStudentTable(); err != nil {
		log.Fatal("Failed to create students table:", err)
	}

}

// Add a new user to the database
func AddUser(username, password, role string) error {
	_, err := DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", username, password, role)
	return err
}

// Get a user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	row := DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Mark attendance for a user
func MarkAttendance(userID, status, date string) error {
	_, err := DB.Exec("INSERT INTO attendance (user_id, status, date) VALUES (?, ?, ?)", userID, status, date)
	return err
}

// Get the attendance report for a specific student or date
func GetAttendanceReport(studentID, date string) ([]Attendance, error) {
	var query string
	var rows *sql.Rows
	var err error

	if studentID != "" && date != "" {
		query = "SELECT user_id, date, status FROM attendance WHERE user_id = ? AND date = ?"
		rows, err = DB.Query(query, studentID, date)
	} else if studentID != "" {
		query = "SELECT user_id, date, status FROM attendance WHERE user_id = ?"
		rows, err = DB.Query(query, studentID)
	} else if date != "" {
		query = "SELECT user_id, date, status FROM attendance WHERE date = ?"
		rows, err = DB.Query(query, date)
	} else {
		query = "SELECT user_id, date, status FROM attendance"
		rows, err = DB.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendance []Attendance
	for rows.Next() {
		var att Attendance
		err := rows.Scan(&att.UserID, &att.Date, &att.Status)
		if err != nil {
			return nil, err
		}
		attendance = append(attendance, att)
	}

	return attendance, nil
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Attendance struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

// ======= Student Management Logic =======

type Student struct {
	ID         int    `json:"id"`
	Roll       string `json:"roll"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Year       string `json:"year"`
}

// Create students table
func InitStudentTable() error {
	createStudentTable := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		roll TEXT UNIQUE,
		name TEXT,
		email TEXT,
		department TEXT,
		year TEXT
	);
	`
	_, err := DB.Exec(createStudentTable)
	return err
}

// Get all students
func GetAllStudents() ([]Student, error) {
	rows, err := DB.Query("SELECT id, roll, name, email, department, year FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		err := rows.Scan(&s.ID, &s.Roll, &s.Name, &s.Email, &s.Department, &s.Year)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

// Add a student
func AddStudent(roll, name, email, department, year string) error {
	query := "INSERT INTO students (roll, name, email, department, year) VALUES (?, ?, ?, ?, ?)"
	_, err := DB.Exec(query, roll, name, email, department, year)
	return err
}

// Update a student
func UpdateStudent(id, roll, name, email, department, year string) error {
	query := "UPDATE students SET roll=?, name=?, email=?, department=?, year=? WHERE id=?"
	_, err := DB.Exec(query, roll, name, email, department, year, id)
	return err
}

// Delete a student
func DeleteStudent(id string) error {
	query := "DELETE FROM students WHERE id=?"
	_, err := DB.Exec(query, id)
	return err
}
