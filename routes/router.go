package routes

import (
	"l-m-s/controllers"

	"github.com/gin-gonic/gin"
)

// A function to register the routes
func Allroutes(router *gin.RouterGroup) {
	// Register routes
	router.POST("/book", controllers.CreateBooK)
	router.GET("/books", controllers.GetBookList)
	router.PUT("/book", controllers.UpdateBook)
	router.DELETE("/book", controllers.DeleteBook)
	router.GET("/book", controllers.GetBookById)
}

// Function SetUp() returns a pointer to a gin.Engine object. This function is used to initialize and configure a web server.
func SetUp() *gin.Engine { //gin.Engine is the main router instance in the Gin web framework, which is used to handle incoming HTTP requests.

	//Creates a new Gin engine instance using the gin.Default() method and assigns it to the router variable.
	router := gin.Default()
	//router.RouterGroup: Represents a group of routes in Gin. It is used to organize and group related routes together.
	Allroutes(&router.RouterGroup)
	return router
}
