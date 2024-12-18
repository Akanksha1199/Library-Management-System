package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AssignBook will assign a book to the student and stored this information in the database
// Parameters accepted:
// book_id:
// student_id:
func IssueBook(c *gin.Context) {

	stuId := c.PostForm("student_id")
	bookId := c.PostForm("book_id")

	studentID, err := strconv.Atoi(stuId)
	if err != nil {
		log.Panicln("[ERROR] controllers.IssueBook - unable to convert string to int for student with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Println("[ERROR] controllers.IssueBook - unable to convert string to int for book with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = models.IssueBook(studentID, bookID)
	if err != nil {
		if err.Error() == "already assigned" {
			c.JSON(http.StatusConflict, gin.H{
				"status": "Failed",
				"error":  "This book is already assigned to someone!"})
		} else {
			//c.JSON(http.StatusBadRequest, gin.H{"message": "No book available"})
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book Issued successfully"})
}

func ReturnBook(c *gin.Context) {

	stuId := c.PostForm("student_id")
	bookId := c.PostForm("book_id")

	log.Printf("studentId- %v BookId - %v", stuId, bookId)

	studentID, err := strconv.Atoi(stuId)
	if err != nil {
		log.Println("[ERROR] controllers.ReturnBook - unable to convert string to int for student with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Println("[ERROR] controllers.ReturnBook - unable to convert string to int for book with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = models.ReturnBook(bookID, studentID)
	if err != nil {
		log.Println("[ERROR] controllers.ReturnBook - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to return book as book is already returned"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Returned successfully"})

}

func ReIssueBook(c *gin.Context) {
	stuId := c.PostForm("student_id")
	bookId := c.PostForm("book_id")
	log.Printf("StudentId- %v BookId - %v", stuId, bookId)

	studentID, err := strconv.Atoi(stuId)
	if err != nil {
		log.Printf("unable to convert string to int %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Printf("unable to convert string to int %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = models.ReIssueBook(bookID, studentID)
	if err != nil {
		log.Println("[ERROR] controllers.ReIssueBook - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to re-issue book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book re-issued Successfully"})
}
