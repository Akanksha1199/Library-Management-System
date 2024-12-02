package main

import (
	"l-m-s/routes"

	_ "github.com/lib/pq"
)

func main() {

	//It sets up the router for further use
	routes := routes.SetUp()

	// Start the server
	routes.Run(":8080")
}
