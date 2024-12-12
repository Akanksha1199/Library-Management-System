package models

import (
	"database/sql"
	"fmt"
	"l-m-s/config"
	"log"
)

// AssignBook
// Accepted Parameters: two variables studentID and bookID of type int
// Returned Parameters: returns an error
func AssignBook(studentID int, bookID int) error {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Printf("error in connecting Database, %v", err)
		return err
	}
	defer db.Close()

	var status sql.NullBool
	query := `SELECT status
	          FROM assign_book
	          WHERE book_id = $1`

	err = db.QueryRow(query, bookID).Scan(&status)
	if err != nil {
		log.Printf("error checking book assignment status: %v", err)
		//return err
	}
	if status.Bool {
		return fmt.Errorf("book with ID %d is already assigned", bookID)
	}

	insertQuery := "INSERT INTO assign_book (student_id, book_id,) VALUES ($1, $2, true)"
	_, err = db.Exec(insertQuery, studentID, bookID)
	if err != nil {
		log.Printf("error inserting assignment record: %v", err)
		return err
	}

	fmt.Println("Book assigned successfully")
	return nil
}
