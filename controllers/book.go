package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetLibraryData handles the GET request to fetch data from the library table.
func GetBookList(c *gin.Context) {
	data, err := models.GetBookList()
	if err != nil {
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

	log.Println("====>", book)

	err = models.CreateBook(book)
	if err != nil {
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

	book, err := models.GetBookById(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func AssignBook(c *gin.Context) {

	stuId := c.PostForm("student_id")
	bookId := c.PostForm("book_id")

	studentID, err := strconv.Atoi(stuId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = models.AssignBook(studentID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Assigned successfully"})

}
