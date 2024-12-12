package controllers

import (
	"l-m-s/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStudent will create new reocrd of student in database with parameters( student_id , student_name, student_email, student_phone, student _dob, student_gender).
// paremeters accepeted in the form of json:
// id: (must be positve integer and not be equal to 0), ERROR:"Student ID must be greater than 0"
// name: must not empty field and should start with an alphabet, ERROR:"unable to create student with:  failed to Create student"
// email: must be unique, ERROR: "unable to create student"
// phone: must be of at-most 12 digits
// dob: must be in a date format
// gender: must be any one out of these (Male, Female or Other)
func CreateStudent(c *gin.Context) {
	var student models.Student

	err := c.BindJSON(&student)
	if err != nil {
		log.Println("unable to bind student data with: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input format"})
		return
	}

	if student.ID <= 0 {
		log.Printf("Student ID must be greater than 0")
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

// GetStudentList fetch students_data from the table, ERROR: "unable to fetch student list"
func GetStudentList(c *gin.Context) {
	data, err := models.GetStudentList()
	if err != nil {
		log.Printf("error fetching student list %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to get student list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// UpdateStudent will update the record of existing table in the database with the given student_id.
// It will update the parameters(student_name, student_email, student_phone, student_dob and student_gender) but not (student_id)
// Parameters accepted:-
// id: (must be positve integer and not be equal to 0), ERROR:"Student ID must be greater than 0"
// name: must not empty field and should start with an alphabet, ERROR:"error updating student with:  failed to Create student"
// email: must be unique, ERROR: "unable to update student"
// phone: must contain atmost 12 digits
// dob: must be in a date format
// gender: must be any one out of these (Male, Female or Other)
func UpdateStudent(c *gin.Context) {
	var student models.Student

	err := c.ShouldBind(&student)
	if err != nil {
		log.Println("error binding request data to student struct with: ", err)
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

// DeleteStudent delete the record of a student from the database with the given student_id
// Parameters accepted: id: (must be positve integer and not be equal to 0), ERROR:"Student ID must be greater than 0"
func DeleteStudent(c *gin.Context) {
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

	err = models.DeleteStudent(studentID)
	if err != nil {
		log.Printf("error deleting student (studentID: %d): %v", studentID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully."})

}

// GetStudentById fetch the record of a student from the database with the given id
// It fetches all the parameters of a student (id, name, email, phone, dob, gender)
// Parameret accepted: id: (must be positve integer and not be equal to 0), ERROR:"Student ID must be greater than 0"
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
