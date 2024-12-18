package config

import (
	"database/sql" //package to interact with database.

	_ "github.com/lib/pq" //pq driver for PostgreSQL

	"fmt" //package for formatting strings.
	"log" //package for logging message in program.
)

// ConnectToDB initializes the connection to the PostgreSQL database.
func ConnectToDB() (*sql.DB, error) {

	// const contains the constant values that never change
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		DBname   = "library"
	)

	//psqlInfo get the String from Sprintf that is a function in Go provided by the fmt package. It is used to format and return a string based on a format specifier and a set of arguments.(format-verbs,arguments)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, DBname)

	// Open the database connection
	/*
		-sql.Open is function used to initialize the databse connection.
		-The first argument ("postgres") specifies the database driver to use (in this case it is PostgreSQL).
		-The second argument (psqlInfo) is a connection string that contains details about how to connect to the database (like username, password, database name, host, and port).
		-DB: A pointer to the sql.DB object that represents the connection pool for the database.
	*/
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Unable to open database: %v\n", err)
		return nil, err
	}

	// Check the connection
	/*
			-This calls the Ping() method on the DB object.
		    -Ping() is used to check if the database connection is working. It sends a small request to the database and expects a response.
		    -If Ping() fails (e.g., the database is unreachable), it returns an error, which gets assigned to the variable err.
	*/
	if err := db.Ping(); err != nil {
		log.Printf("Unable to connect to the database: %v\n", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
