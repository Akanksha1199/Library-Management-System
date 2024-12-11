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

// CREATE TABLE student(
// 	id SERIAL PRIMARY KEY,
// 	name varchar(50) NOT NULL,
// 	email VARCHAR(100) UNIQUE NOT NULL,
// 	phone VARCHAR(12) NOT NULL CHECK(phone IN('+[0,9]')),
// 	dob DATE,
// 	gender VARCHAR(15) CHECK(gender IN('Male', 'Female', 'Other')),
// 	enrollment_date DATE
// 	);
