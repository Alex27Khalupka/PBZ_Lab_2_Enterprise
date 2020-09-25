package service

import (
	"database/sql"
)

func DeleteEmployee(db *sql.DB, employeeID string) error {
	if err := DeleteMovementOfEmployee(db, employeeID); err !=nil{
		return err
	}
	query := "DELETE FROM employees WHERE employee_number = $1"
	stmt, err := db.Prepare(query)
	if err !=nil{
		return err
	}
	_, err = stmt.Exec(employeeID)
	return err
}

func DeleteMovementOfEmployee(db *sql.DB, employeeID string) error {

	query := "DELETE FROM movement_of_employees WHERE employee_number = $1"
	stmt, err := db.Prepare(query)
	if err !=nil{
		return err
	}
	_, err = stmt.Exec(employeeID)

	return nil
}

