package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"time"
)

func POSTEmployee(db *sql.DB, employee model.Employee, divisionID string) error{

	query := "INSERT INTO employees (employee_number, first_name, last_name, second_name, position, age, sex) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
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
	if _, err := db.Query("INSERT INTO movement_of_employees (employee_number, movement_date, division_number) VALUES" +
	" (?, ?, ?)", employeeNumber, time.Now().Format("2020-01-01"), divisionID); err!=nil {
		return err
	}
	return nil
}
