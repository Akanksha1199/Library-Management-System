package controllers

import (
	"l-m-s/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// This function handles http request to assign a book to a student.
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
