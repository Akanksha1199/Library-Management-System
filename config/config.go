package config

import (
	"database/sql" //package to interact with database.
	"fmt"          //package for formatting strings.
	"log"          //package for logging message in program.
)

// *sql.DB is a pointer to a sql.DB object provided by Go's database/sql package. The sql.DB type represents the database connection pool
var db *sql.DB

// ConnectToDB initializes the connection to the PostgreSQL database.
func ConnectToDB() error {

	// const contains the constant values that never change
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1234"
		dbname   = "library"
	)

	//psqlInfo get the String from Sprintf that is a function in Go provided by the fmt package. It is used to format and return a string based on a format specifier and a set of arguments.(format-verbs,arguments)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Here error is a built-in Interface in Go that is used to represent errors. It can store values when something went wrong.
	var err error
	/*
		-sql.Open is function used to initialize the databse connection.
		-The first argument ("postgres") specifies the database driver to use (in this case it is PostgreSQL).
		-The second argument (psqlInfo) is a connection string that contains details about how to connect to the database (like username, password, database name, host, and port).
		-db: A pointer to the sql.DB object that represents the connection pool for the database.
	*/
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Unable to open database: %v\n", err)
		return err
	}

	/*
			-This calls the Ping() method on the db object.
		    -Ping() is used to check if the database connection is working. It sends a small request to the database and expects a response.
		    -If Ping() fails (e.g., the database is unreachable), it returns an error, which gets assigned to the variable err.
	*/
	if err := db.Ping(); err != nil {
		log.Printf("Unable to connect to the database: %v\n", err)
		return err
	}

	//If everything works fine, then db connected successfully.
	fmt.Println("Successfully connected to PostgreSQL")
	return nil
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
