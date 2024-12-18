package models

import (
	"fmt"
	"l-m-s/config"
	"log"
)

func ApplyFine(bookID int) error {
	db, err := config.ConnectToDB()
	if err != nil {
		return fmt.Errorf("error in connecting database: %v", err)
	}
	defer db.Close()

	query := `SELECT 
    student_id,
    book_id,
    issue_date,
    returned_at,
    GREATEST(0, EXTRACT(DAY FROM (COALESCE(returned_at, CURRENT_DATE) - issue_date)) - 7) AS overdue_days,
    GREATEST(0, EXTRACT(DAY FROM (COALESCE(returned_at, CURRENT_DATE) - issue_date)) - 7) * 10 AS fine
FROM 
    assign_book`

	log.Println(query)

	_, err = db.Exec(query)
	if err != nil {
		log.Println("error checking fine with: ", err)
	}
	fmt.Println("Fine is applied successfully")
	return nil
}
