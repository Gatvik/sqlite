package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func getDb() *sql.DB {
	ensureCreation()

	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Panicf("Can't connect to DB: %s", err)
	}

	createTableIfNotExists(db)
	return db
}

func ensureCreation() {
	if file, err := os.Open("students.db"); err != nil {
		os.Create("students.db")
		file.Close()
	}
}

func createTableIfNotExists(db *sql.DB) {
	var query = `
	CREATE TABLE IF NOT EXISTS students(
	  Id INTEGER PRIMARY KEY AUTOINCREMENT,
	  Name TEXT NOT NULL,
	  Phone TEXT NOT NULL 
	);
`
	_, err := db.Exec(query)
	if err != nil {
		log.Panicf("Can't create table: %s", err)
	}
}

func addNewStudent(db *sql.DB, name string, phone string) {
	var query = `INSERT INTO students (Name, Phone) VALUES (?, ?)`
	_, err := db.Exec(query, name, phone)
	if err != nil {
		log.Printf("Can't add new student: %s", err)
	}
}

func getStudents(db *sql.DB) []Student {
	data, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Panicf("Can't get data from db: %s", err)
	}
	defer data.Close()
	var students []Student

	for data.Next() {
		var student = Student{}
		err := data.Scan(&student.Id, &student.Name, &student.Phone)
		if err != nil {
			log.Panicf("Can't scan data: %s", err)
		}

		students = append(students, student)
	}
	return students
}

func deleteStudent(db *sql.DB, id int) {
	var query = `DELETE FROM students WHERE Id = ?`

	_, err := db.Exec(query, id)
	if err != nil {
		log.Panicf("Can't delete student: %s", err)
	}
}
