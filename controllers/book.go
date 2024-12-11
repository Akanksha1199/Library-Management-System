package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBookList handles the GET request to fetch data from the book table.
func GetBookList(c *gin.Context) {
	data, err := models.GetBookList()
	if err != nil {
		log.Printf("unable to fetch book list %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

/*
- c is the context object provided by the Gin framework, which allows you to handle HTTP requests and responses.
*/
func CreateBooK(c *gin.Context) {
	var book models.Book

	//This takes data sent by the user in the request (e.g., JSON or form data) and tries to fill the book variable with it.
	err := c.ShouldBind(&book)
	if err != nil {
		log.Printf("unable to bind book details %v", err)
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

	log.Println("====>", book)

	err = models.CreateBook(book)
	if err != nil {
		log.Printf("unable to create book %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBind(&book)
	if err != nil {
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
