package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"log"
	"time"
)

func POSTEmployee(db *sql.DB, employee model.Employee, divisionID string) error{

	query := "INSERT INTO employees (employee_number, first_name, last_name, second_name, position, age, sex) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("kek ", query)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(employee.EmployeeNumber, employee.FirstName, employee.LastName, employee.SecondName, employee.Position,
		employee.Age, employee.Sex)

	if err !=nil{
		return err
	}

	//if _, err := db.Query("INSERT INTO employees (employee_number, first_name, last_name, " +
	//	"second_name, position, age, sex) VALUES (?, ?, ?, ?, ?, ?, ?)",
	//	employee.EmployeeNumber, employee.FirstName, employee.LastName, employee.SecondName, employee.Position,
	//	employee.Age, employee.Sex); err!=nil{
	//	return err
	//}

	err = POSTMovementOfEmployees(db, employee.EmployeeNumber, divisionID)

	return nil
}

func POSTMovementOfEmployees(db *sql.DB, employeeNumber string, divisionID string) error{
	query := "INSERT INTO movement_of_employees (employee_number, movement_date, division_number) VALUES" +
		" ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(employeeNumber, time.Now().Format("2020-01-01"), divisionID)

	if err !=nil{
		return err
	}

	return nil
}
