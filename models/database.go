package models

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// create a book struct
type Book struct {
	ID       int    `json:"id"` //These json tag tells Go's JSON library how to map this field when converting the struct to/from JSON.
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Returned bool   `json:"returned"`
}

/*
This function fetch or return a list of books.
The function returns a slice of Book structs. The function also returns an error type.
*/
func GetBookList() ([]Book, error) {
	//.Query is a method used to run an SQL query against the database. It retrieves data from the database based on the SQL command provided.
	rows, err := db.Query("SELECT id, name, cost FROM book")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	//defer is a keyword in Go that schedules a function to be executed later (after the surrounding function completes).
	//rows.Close() is a method that closes the rows returned by a database query.
	defer rows.Close()

	//[]Book means you are declaring a slice that will hold multiple Book objects.
	var results []Book
	//In this for loop, rows is the result of a query to the database, and Next() is a method that moves the cursor to the next row in the result set.
	for rows.Next() {
		var id, cost int
		var name string
		if err := rows.Scan(&id, &name, &cost); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		//It(results) starts as an empty slice, and new books are added to it using the append() function.
		results = append(results, Book{
			ID:   id,
			Name: name,
			Cost: cost,
		})
	}

	return results, nil
}

/*
It is a function named as CreateBook.
-book is a parameter passed into the function.
-Book is the type of the parameter, meaning the function expects an argument of type Book.
*/
func CreateBook(book Book) error {

	//query is a variable that holds the SQL query as a string.
	query := "INSERT INTO book (id, name, cost ) VALUES ($1, $2, $3)"

	//_: The underscore is used to ignore the result of a function when it's not needed.
	//db.Exec is a method from the *sql.DB type in Go that is used to execute a query on the database.
	_, err := db.Exec(query, book.ID, book.Name, book.Cost)
	if err != nil {
		return fmt.Errorf("failed to Create book: %v", err)
	}

	fmt.Println("Book created successfully")
	return nil
}

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

		if book.Cost <= 0 {
			return fmt.Errorf("cost must be greater than 0")
		}
		query := "UPDATE book SET cost = $1 WHERE id = $2"
		_, err := db.Exec(query, book.Cost, book.ID)
		if err != nil {
			return fmt.Errorf("failed to Update book: %v", err)
		}

	case 3:
		if book.Cost <= 0 {
			return fmt.Errorf("cost must be greater than 0")
		}
		query := "UPDATE book SET name = $1, cost = $2 WHERE id = $3"
		_, err := db.Exec(query, book.Name, book.Cost, book.ID)
		if err != nil {
			return fmt.Errorf("failed to Update book: %v", err)
		}
	}

	fmt.Println("Book Updated successfully")
	return nil

}

// This function will use this bookID to identify which book to delete from the database or list.
func DeleteBook(bookID int) error {

	query := "DELETE FROM book WHERE id = $1"

	_, err := db.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("failed to Delete book: %v", err)
	}

	fmt.Println("Book Deleted successfully")
	return nil
}

// Book is the return type of the function. The function will return a Book struct and error
func GetBookById(id int) (Book, error) {
	var cost int
	var name string
	var book Book

	query := "SELECT id,name,cost FROM book WHERE id = $1"

	//QueryRow(query, id): executes the SQL query but expects only one row to be returned.
	//.Scan(&id, &name, &cost) : is a method that reads the data returned by QueryRow and places it into the variables id, name, and cost.
	err := db.QueryRow(query, id).Scan(&id, &name, &cost)
	if err != nil {
		log.Printf("no book found with id %d", id)
		return book, err
	}
	//book = Book{}: This line creates a new instance of the Book struct.
	book = Book{
		ID:   id,
		Name: name,
		Cost: cost,
	}
	return book, nil
}
