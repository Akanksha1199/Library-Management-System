package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student

	err := c.BindJSON(&student)
	if err != nil {
		log.Printf("unable to bind student data %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input format"})
		return
	}

	if student.ID <= 0 {
		log.Println("Student ID must be greater than 0")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student ID must be greater than 0"})
		return
	}

	err = models.CreateStudent(student)
	if err != nil {
		log.Printf("error creating student (studentID: %d): %v", student.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student})
}

func GetStudentList(c *gin.Context) {
	data, err := models.GetStudentList()
	if err != nil {
		log.Printf("error fetching student list %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to get student list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student

	err := c.ShouldBind(&student)
	if err != nil {
		log.Printf("error binding request data to student struct: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if student.ID <= 0 {
		log.Println("Student ID must be a positive number")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student ID must be a positive number"})
		return
	}

	err = models.UpdateStudent(student)
	if err != nil {
		log.Printf("error updating student(studentID: %d) %v", student.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func DeleteStudent(c *gin.Context) {
	id := c.Query("id")
	studentID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("unable to convert string to int %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if studentID <= 0 { //Q- Why here not student.ID??
		log.Println("Student ID must be a positive number")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student ID must be a positive number"})
		return
	}

	err = models.DeleteStudent(studentID)
	if err != nil {
		log.Printf("error deleting student (studentID: %d): %v", studentID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully."})

}

func GetStudentById(c *gin.Context) {

	id := c.Query("id")

	studentID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("unable to convert string to int %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if studentID <= 0 {
		log.Println("Student ID must be a positive number")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Student ID must be a positive number"})
		return
	}

	student, err := models.GetStudentById(studentID)
	if err != nil {
		log.Printf("unable to get student by ID (studentID: %d): %v", studentID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}
