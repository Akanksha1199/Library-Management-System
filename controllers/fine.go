package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApplyFine(c *gin.Context) {
	stuId := c.PostForm("student_id")
	bookId := c.PostForm("book_id")
	log.Printf("StudentId- %v BookId - %v", stuId, bookId)

	bookID, err := strconv.Atoi(bookId)
	if err != nil {
		log.Printf("unable to convert string to int %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	err = models.ApplyFine(bookID)
	if err != nil {
		log.Println("[ERROR] controllers.ApplyFine - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to apply fine to the student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fine Applied successfully"})
}
