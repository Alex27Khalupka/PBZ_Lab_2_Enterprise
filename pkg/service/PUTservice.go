package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
)

func PUTInventory(db *sql.DB, inventoryToUpdate model.Inventory, inventoryID string) (model.Inventory, error) {

	inventory := GetInventoryByID(db, inventoryID)

	if len(inventoryToUpdate.InventoryNumber) < 1 {
		inventoryToUpdate.InventoryNumber = inventory.InventoryNumber
	}

	if len(inventoryToUpdate.InventoryName) < 1 {
		inventoryToUpdate.InventoryName = inventory.InventoryName
	}

	if len(inventoryToUpdate.InventoryModel) < 1 {
		inventoryToUpdate.InventoryModel = inventory.InventoryModel
	}

	if inventoryToUpdate.YearOfIssue.IsZero() {
		inventoryToUpdate.YearOfIssue = inventory.YearOfIssue
	}

	if err := UpdateInventory(db, inventoryToUpdate, inventoryID); err != nil {
		return inventory, err
	}

	return inventoryToUpdate, nil
}

func UpdateInventory(db *sql.DB, inventoryToUpdate model.Inventory, inventoryID string) error {
	query := "UPDATE inventory SET inventory_number = $1, inventory_name = $2, inventory_model = $3, " +
		"year_of_issue = $4 WHERE inventory_number = $5"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventoryToUpdate.InventoryNumber, inventoryToUpdate.InventoryName,
		inventoryToUpdate.InventoryModel, inventoryToUpdate.YearOfIssue, inventoryID)

	return err
}

func PUTEmployee(db *sql.DB, employeeToUpdate model.Employee, employeeID string) (model.Employee, error) {

	employee := GetEmployeeByID(db, employeeID)

	if len(employeeToUpdate.EmployeeNumber) < 1 {
		employeeToUpdate.EmployeeNumber = employee.EmployeeNumber
	}

	if len(employeeToUpdate.FirstName) < 1 {
		employeeToUpdate.FirstName = employee.FirstName
	}

	if len(employeeToUpdate.LastName) < 1 {
		employeeToUpdate.LastName = employee.LastName
	}

	if len(employeeToUpdate.SecondName) < 1 {
		employeeToUpdate.SecondName = employee.SecondName
	}

	if len(employeeToUpdate.Position) < 1 {
		employeeToUpdate.Position = employee.Position
	}

	if employeeToUpdate.Age == 0 {
		employeeToUpdate.Age = employee.Age
	}

	if len(employeeToUpdate.Sex) < 1 {
		employeeToUpdate.Sex = employee.Sex
	}

	if err := UpdateEmployee(db, employeeToUpdate, employeeID); err != nil {
		return employee, err
	}

	return employeeToUpdate, nil
}

func UpdateEmployee(db *sql.DB, employeeToUpdate model.Employee, employeeID string) error {
	query := "UPDATE employees SET employee_number = $1, first_name = $2, last_name = $3, second_name = $4, " +
		"position = $5, age = $6, sex = $7 WHERE employee_number = $8"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(employeeToUpdate.EmployeeNumber, employeeToUpdate.FirstName, employeeToUpdate.LastName,
		employeeToUpdate.SecondName, employeeToUpdate.Position, employeeToUpdate.Age, employeeToUpdate.Sex, employeeID)

	return nil
}
