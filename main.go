package main

import (
	"l-m-s/config"
	"l-m-s/routes"

	_ "github.com/lib/pq"
)

func main() {

	// Connect to the database
	config.ConnectToDB()

	// Close the databse
	defer config.CloseDB()

	//It sets up the router for further use
	routes := routes.SetUp()

	// Start the server
	routes.Run(":8080")

}
