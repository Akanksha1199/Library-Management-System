package models

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// ConnectToDB initializes the connection to the PostgreSQL database.
func ConnectToDB() error {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "library"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Unable to open database: %v\n", err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Unable to connect to the database: %v\n", err)
		return err
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return nil
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// create a book struct
type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Returned bool   `json:"returned"`
}

func GetBookList() ([]Book, error) {
	rows, err := db.Query("SELECT id, name , cost FROM book")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var results []Book
	for rows.Next() {
		var id, cost int
		var name string
		if err := rows.Scan(&id, &name, &cost); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		// Collect results into a map--------- Map Approach is changed
		results = append(results, Book{
			ID:   id,
			Name: name,
			Cost: cost,
		})
	}

	return results, nil
}

func CreateBook(book Book) error {
	query := "INSERT INTO book (id, name, cost) VALUES ($1, $2, $3)"

	_, err := db.Exec(query, book.ID, book.Name, book.Cost)
	if err != nil {
		return fmt.Errorf("failed to Create book: %v", err)
	}

	fmt.Println("Book created successfully")
	return nil
}

// UPDATE book
// SET book.Name = 'java2.0'
// WHERE book.id = 1;

// UPDATE book
// SET name = 'Ruby2.0', cost = '500'
// WHERE id = 5

func UpdateBook(book Book) error {
	//query := "UPDATE book SET name = $1, cost = $2 WHERE id = $3"

	fmt.Println("What do you want to update...?")
	fmt.Println("1. Book Name ")
	fmt.Println("2. Book Cost")
	fmt.Println("3. Both Book Name and Cost")

	var choose int
	fmt.Scanln(&choose)

	switch choose {

	case 1:

		query := "UPDATE book SET name = $1 WHERE id = $2"
		_, err := db.Exec(query, book.Name, book.ID)
		if err != nil {
			return fmt.Errorf("failed to Update book: %v", err)
		}

	case 2:

		query := "UPDATE book SET cost = $1 WHERE id = $2"
		_, err := db.Exec(query, book.Cost, book.ID)
		if err != nil {
			return fmt.Errorf("failed to Update book: %v", err)
		}

	case 3:
		query := "UPDATE book SET name = $1, cost = $2 WHERE id = $3"
		_, err := db.Exec(query, book.Name, book.Cost, book.ID)
		if err != nil {
			return fmt.Errorf("failed to Update book: %v", err)
		}
	}

	// 	_, err := db.Exec(query, book.Name, book.Cost, book.ID)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to Update book: %v", err)
	// 	}

	fmt.Println("Book Updated successfully")
	return nil

}

func DeleteBook(bookID int) error {

	query := "DELETE FROM book WHERE id = $1"

	_, err := db.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("failed to Delete book: %v", err)
	}

	fmt.Println("Book Deleted successfully")
	return nil
}

func GetBookById(id int) (Book, error) {
	var cost int
	var name string
	var book Book

	query := "SELECT id,name,cost FROM book WHERE id = $1"

	err := db.QueryRow(query, id).Scan(&id, &name, &cost)
	if err != nil {
		log.Printf("no book found with id %d", id)
		return book, err
	}
	book = Book{
		ID:   id,
		Name: name,
		Cost: cost,
	}
	return book, nil
}
