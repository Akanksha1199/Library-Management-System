package models

import (
	"fmt"
	"l-m-s/config"
	"log"
)

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DOB        string `json:"dob"`
	Gender     string `json:"gender"`
	Created_At string `json:"created_at,omitempty"`
}

func CreateStudent(student Student) error {
	db, err := config.ConnectToDB()
	if err != nil {
		return fmt.Errorf("unable to connect to database, %v", err)
	}
	defer db.Close()

	query := `INSERT INTO student(id, name, email, phone, dob, gender, created_at)
              VALUES ($1,$2,$3,$4,$5,$6,NOW())`

	_, err = db.Exec(query, student.ID, student.Name, student.Email, student.Phone, student.DOB, student.Gender)

	if err != nil {
		return fmt.Errorf("failed to create student,%v", err)
	}
	fmt.Println("Student created successfully")
	return nil

}

func GetStudentList() ([]Student, error) {
	db, err := config.ConnectToDB()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database %v", err)
	}
	defer db.Close()

	query := `SELECT id, name, email, phone, dob, gender, created_at
                           FROM student;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed %v", err)
	}
	defer rows.Close()

	var result []Student

	var id int
	var name, email, phone, dob, gender, created_at string
	for rows.Next() {

		if err := rows.Scan(&id, &name, &email, &phone, &dob, &gender, &created_at); err != nil {
			return nil, fmt.Errorf("unable to scan row %v", err)
		}
		result = append(result, Student{
			ID:         id,
			Name:       name,
			Email:      email,
			Phone:      phone,
			DOB:        dob,
			Gender:     gender,
			Created_At: created_at,
		})
	}
	return result, nil
}

func UpdateStudent(student Student) error {
	var SetValues string

	db, err := config.ConnectToDB()
	if err != nil {
		return fmt.Errorf("error in connecting database: %v", err)
	}

	defer db.Close()

	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	if len(student.Name) > 0 {
		SetValues += " name = '" + student.Name + "'"
	}
	if len(student.Email) > 0 {
		SetValues += fmt.Sprintf(" ,email ='"+"%v"+"'", student.Email)
	}
	if len(student.Phone) > 0 {
		SetValues += fmt.Sprintf(" ,phone = '"+"%v"+"'", student.Phone)
	}
	if len(student.DOB) > 0 {
		SetValues += fmt.Sprintf(" ,dob = '"+"%s"+"'", student.DOB)
	}
	if len(student.Gender) > 0 {
		SetValues += fmt.Sprintf(" ,gender = '"+"%v'", student.Gender)
	}

	if len(SetValues) == 0 {
		fmt.Println("There is no value to Update")
		return nil
	}

	query := `UPDATE student
	SET ` + SetValues +
		`,updated_at=NOW() 
		WHERE id = $1`

	log.Println(query)

	_, err = db.Exec(query, student.ID)
	if err != nil {
		return fmt.Errorf("failed to Update student: %v", err)
	}

	fmt.Println("Student Updated successfully")
	return nil

}

func DeleteStudent(studentID int) error {
	db, err := config.ConnectToDB()
	if err != nil {
		return fmt.Errorf("error in connecting database: %v", err)
	}
	defer db.Close()

	query := "DELETE FROM student WHERE id = $1"

	_, err = db.Exec(query, studentID)
	if err != nil {
		return fmt.Errorf("failed to Delete student: %v", err)

	}
	fmt.Println("Student Deleted successfully")
	return nil
}

func GetStudentById(id int) (Student, error) {
	db, err := config.ConnectToDB()
	if err != nil {
		return Student{}, fmt.Errorf("error in connecting database: %v", err)
	}
	defer db.Close()
	var name, email, phone, dob, gender, created_at string
	var student Student

	query := `SELECT id, name, email, phone, dob, gender, created_at
              FROM student
              WHERE id = $1`

	err = db.QueryRow(query, id).Scan(&id, &name, &email, &phone, &dob, &gender, &created_at)
	if err != nil {
		log.Printf("no student found with id %d", id)
		return student, err
	}

	student = Student{
		ID:         id,
		Name:       name,
		Email:      email,
		Phone:      phone,
		DOB:        dob,
		Gender:     gender,
		Created_At: created_at,
	}
	return student, nil
}
