package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"log"
)

func GetDivisions(db *sql.DB) []model.Division {

	rows, err := db.Query("SELECT divisions.division_number, divisions.division_name FROM divisions")
	if err != nil {
		log.Fatal(err)
	}

	var divisions []model.Division
	for rows.Next() {
		var  division model.Division
		if err := rows.Scan(&division.DivisionNumber, &division.DivisionName); err != nil {
			log.Fatal(err)
		}
		divisions = append(divisions, division)
	}

	return divisions
}

func GetEmployees(db *sql.DB) []model.Employee {
	rows, err := db.Query("SELECT employees.employee_number, employees.first_name, employees.last_name, employees.second_name, employees.position, employees.age, employees.sex FROM employees")
	if err != nil {
		log.Fatal(err)
	}

	var employees []model.Employee
	for rows.Next() {
		var  employee model.Employee
		if err := rows.Scan(&employee.EmployeeNumber, &employee.FirstName, &employee.LastName, &employee.SecondName, &employee.Position, &employee.Age, &employee.Sex); err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}

	return employees
}