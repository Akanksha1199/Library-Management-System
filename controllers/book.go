package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBookList fetch books_data from the table, ERROR: "unable to fetch book list"
func GetBookList(c *gin.Context) {
	data, err := models.GetBookList()
	if err != nil {
		log.Printf("unable to fetch book list %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// CreateBook will create new reocrd of book in database with parameters( bookid , bookname, cost of the book).
// paremeters accepeted in the form of json:
// id: (must be positve integer and not be equal to 0), ERROR:"Book ID must be greater than 0"
// name: must not empty field and should start with an alphabet, ERROR:"unable to create book with:  failed to Create book"
// cost: must be greater than 0, ERROR:"cost should not be less than or equals to zero"
func CreateBooK(c *gin.Context) {
	var book models.Book

	//This takes data sent by the user in the request (e.g., JSON or form data) and tries to fill the book variable with it.
	err := c.ShouldBind(&book)
	if err != nil {
		log.Println("unable to bind book details with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{ //gin.H: This is a shortcut provided by Gin for creating a map[string]interface{}.

			"error": err.Error(),
		})
		return
	}
	if book.Cost <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cost should not be less than or equals to zero",
		})
		return
	}

	if book.ID <= 0 {
		log.Println("Book ID must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book ID must be greater than 0"})
		return
	}

	err = models.CreateBook(book)
	if err != nil {
		log.Println("unable to create book with: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook will update the record of existing table in the database with the given book_id.
// It will update the parameters(book_name and book_cost) but not (book_id)
// Parameters accepted:-
// id: (must be positve integer and not be equal to 0), ERROR:"Book ID must be greater than 0"
// name: must not empty field and should start with an alphabet, ERROR:"unable to create book with:  failed to Create book"
// cost: must be greater than 0, ERROR:"cost should not be less than or equals to zero"
func UpdateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBind(&book)
	if err != nil {
		log.Println("error binding request data to student struct with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if book.ID <= 0 {
		log.Println("Book ID must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book ID must be greater than 0"})
		return
	}

	err = models.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook delete the record of a book from the database with the given book_id
// Parameters accepted: id: (must be positve integer and not be equal to 0), ERROR:"Book ID must be greater than 0"
func DeleteBook(c *gin.Context) {
	id := c.Query("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if bookID <= 0 {
		log.Println("Book ID must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book ID must be greater than 0"})
		return
	}

	err = models.DeleteBook(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully."})

}

// GetBookById fetch the record of a book from the database with the given id
// It fetches all the parameters of a book (book_id, book_name, book_cost)
// Parameret accepted: id: (must be positve integer and not be equal to 0), ERROR:"Book ID must be greater than 0"
func GetBookById(c *gin.Context) {

	id := c.Query("id")

	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if bookID <= 0 {
		log.Println("Book ID must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book ID must be greater than 0"})
		return
	}

	book, err := models.GetBookById(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
