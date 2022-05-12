package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Creating db and table in it, getting a pointer to sql.DB structure
	var db = getDb()
	defer db.Close()
	//Adding new students to db
	fmt.Println("Adding new students")
	addNewStudent(db, "Ivan", "8-800-555-35-35")
	addNewStudent(db, "Dima", "348943984389")
	//Deleting student from db by id
	deleteStudent(db, 4)
	//Getting all students from db
	var students = getStudents(db)
	fmt.Println("Displaying students")
	fmt.Println(students)
}
