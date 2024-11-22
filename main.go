package main

import (
	"l-m-s/controllers"
	"l-m-s/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// Connect to the database
	models.ConnectToDB()
	defer models.CloseDB()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	router.POST("/book", controllers.CreateBooK)
	router.GET("/books", controllers.GetBookList)

	// Start the server
	router.Run(":8080")
}
