package main

import (
	"log"
	"os"

	"l-m-s/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Get the port from the .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
		log.Println("PORT not specified in .env file or environment. Using default:", port)
	}

	// Set up the router for further use
	r := routes.SetUp()

	// Start the server
	log.Println("Starting server on port:", port)
	r.Run(":" + port)
}
