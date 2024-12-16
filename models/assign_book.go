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
func IssueBook(studentID int, bookID int) error {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Printf("error in connecting Database, %v", err)
		return err
	}
	defer db.Close()

	var status sql.NullBool
	query := `SELECT status
	          FROM assign_book
	          WHERE status = true AND
			  book_id = $1`

	err = db.QueryRow(query, bookID).Scan(&status)
	if err != nil {
		log.Println("error checking book assignment status with: ", err)
	}
	if status.Bool {
		log.Printf("book with ID %d is already assigned %v", bookID, err)
		return fmt.Errorf("already assigned")
	}

	insertQuery := `INSERT INTO assign_book (student_id, book_id, status) 
	                VALUES ($1, $2, true)`
	_, err = db.Exec(insertQuery, studentID, bookID)
	if err != nil {
		log.Println("error inserting assignment record with: ", err)
		return fmt.Errorf("no such book available")
	}

	fmt.Println("Book assigned successfully")
	return nil
}

func ReturnBook(bookID int, studentID int) error {

	db, err := config.ConnectToDB()
	if err != nil {
		log.Println("error in connecting Database with: ", err)
		return err
	}
	defer db.Close()

	var status bool
	query := `SELECT status
	          FROM assign_book
	          WHERE status = true AND
			  book_id = $1`

	err = db.QueryRow(query, bookID).Scan(&status)
	if err != nil {
		log.Println("error checking book assignment status with: ", err)
	}

	if !status {
		log.Println("book is already returned and cannot be returned again ", err)
		return err
	}

	query = `UPDATE assign_book
	          SET status = false, returned_at = NOW()
			  WHERE book_id = $1 AND student_id = $2
			  RETURNING id`

	log.Println(query)

	var id int
	err = db.QueryRow(query, bookID, studentID).Scan(&id)
	if err != nil {
		log.Println("[ERROR] models.ReturnBook - query failed with: ", err)
		return err
	}
	fmt.Println("Book is returned successfully")
	return nil
}

func ReIssueBook(bookID int, studentID int) error {
	db, err := config.ConnectToDB()
	if err != nil {
		log.Println("error in connecting Database with: ", err)
		return err
	}
	defer db.Close()

	var status bool
	query := `SELECT status
	          FROM assign_book
	          WHERE status = true AND
			  book_id = $1`

	err = db.QueryRow(query, bookID).Scan(&status)
	if err != nil {
		log.Println("error checking book assignment status with: ", err)
	}

	if status {
		return fmt.Errorf("book is already issued ")
	}

	query = `UPDATE assign_book
	          SET status = true, issue_date = NOW()
			  WHERE book_id = $1 AND student_id = $2
			  RETURNING id`

	log.Println(query, bookID, studentID)

	var id int
	err = db.QueryRow(query, bookID, studentID).Scan(&id)
	if err != nil {
		log.Println("[ERROR] models.ReIssueBook - query failed with: ", err)
		return err
	}
	fmt.Println("Book re-issued successfully")
	return nil
}
