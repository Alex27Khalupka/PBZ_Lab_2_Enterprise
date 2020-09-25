package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"log"
	"time"
)

func GetDivisions(db *sql.DB) []model.Division {

	rows, err := db.Query("SELECT divisions.division_number, divisions.division_name FROM divisions")
	if err != nil {
		log.Fatal(err)
	}

	var divisions []model.Division
	for rows.Next() {
		var division model.Division
		if err := rows.Scan(&division.DivisionNumber, &division.DivisionName); err != nil {
			log.Fatal(err)
		}
		divisions = append(divisions, division)
	}

	return divisions
}

func GetEmployees(db *sql.DB) []model.Employee {
	rows, err := db.Query("SELECT employees.employee_number, employees.first_name, employees.last_name, " +
		"employees.second_name, employees.position, employees.age, employees.sex FROM employees")
	if err != nil {
		log.Fatal(err)
	}

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.EmployeeNumber, &employee.FirstName, &employee.LastName, &employee.SecondName,
			&employee.Position, &employee.Age, &employee.Sex); err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}

	return employees
}

func GetInventory(db *sql.DB) []model.Inventory {
	rows, err := db.Query("SELECT inventory.inventory_number, inventory_name, inventory.inventory_model, " +
		"inventory.year_of_issue FROM inventory")
	if err != nil {
		log.Fatal(err)
	}

	var inventoryList []model.Inventory
	for rows.Next() {
		var inventory model.Inventory
		if err := rows.Scan(&inventory.InventoryNumber, &inventory.InventoryName, &inventory.InventoryModel,
			&inventory.YearOfIssue); err != nil {
			log.Fatal(err)
		}
		inventoryList = append(inventoryList, inventory)
	}

	return inventoryList
}

func GetRepairs(db *sql.DB) []model.Repair {
	rows, err := db.Query("SELECT repairs.repair_id, repairs.inventory_number, repairs.service_start_date, " +
		"repairs.repair_type, repairs.days_to_repair, repairs.employee_number, repairs.waybill_number FROM repairs")

	if err != nil {
		log.Fatal(err)
	}

	var repairsList []model.Repair
	for rows.Next() {
		var repair model.Repair
		if err := rows.Scan(&repair.RepairID, &repair.InventoryNumber, &repair.ServiceStartDay, &repair.RepairType,
			&repair.RepairTime, &repair.EmployeeNumber, &repair.WaybillNumber); err != nil {
			log.Fatal(err)
		}
		repairsList = append(repairsList, repair)
	}

	return repairsList
}

func GetWaybills(db *sql.DB) []model.Waybill {
	rows, err := db.Query("SELECT waybills.waybill_number, waybills.receiving_date, waybills.price, " +
		"waybills.detail_name FROM waybills")

	if err != nil {
		log.Fatal(err)
	}

	var waybillsList []model.Waybill
	for rows.Next() {
		var waybill model.Waybill
		if err := rows.Scan(&waybill.WaybillNumber, &waybill.ReceivingDate, &waybill.Price,
			&waybill.DetailName); err != nil {
			log.Fatal(err)
		}
		waybillsList = append(waybillsList, waybill)
	}
	return waybillsList
}

func GetMovementOfEmployees(db *sql.DB) []model.MovementOfEmployees {
	rows, err := db.Query("SELECT movement_of_employees.employee_number, movement_of_employees.movement_date, " +
		"movement_of_employees.division_number FROM movement_of_employees")

	if err != nil {
		log.Fatal(err)
	}

	var movementOfEmployeesList []model.MovementOfEmployees
	for rows.Next() {
		var movementOfEmployees model.MovementOfEmployees
		if err := rows.Scan(&movementOfEmployees.EmployeeNumber, &movementOfEmployees.MovementDate,
			&movementOfEmployees.DivisionNumber); err != nil {
			log.Fatal(err)
		}
		movementOfEmployeesList = append(movementOfEmployeesList, movementOfEmployees)
	}
	return movementOfEmployeesList
}

func GetMovementOfInventory(db *sql.DB) []model.MovementOfInventory {
	rows, err := db.Query("SELECT movement_of_inventory.inventory_number, movement_of_inventory.movement_date, " +
		"movement_of_inventory.division_number FROM movement_of_inventory")

	if err != nil {
		log.Fatal(err)
	}

	var movementOfInventoryList []model.MovementOfInventory
	for rows.Next() {
		var movementOfInventory model.MovementOfInventory
		if err := rows.Scan(&movementOfInventory.InventoryNumber, &movementOfInventory.MovementDate,
			&movementOfInventory.DivisionNumber); err != nil {
			log.Fatal(err)
		}
		movementOfInventoryList = append(movementOfInventoryList, movementOfInventory)
	}
	return movementOfInventoryList
}

func GetEmployeesByDivision(db *sql.DB, id string) []model.EmployeeResponse {
	rows, err := db.Query("SELECT employees.first_name, employees.last_name, employees.second_name, " +
		"employees.age FROM employees INNER JOIN movement_of_employees ON employees.employee_number = " +
		"movement_of_employees.employee_number WHERE movement_of_employees.division_number = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	var employees []model.EmployeeResponse
	for rows.Next() {
		var employee model.EmployeeResponse
		var age int64
		if err := rows.Scan(&employee.FirstName, &employee.LastName, &employee.SecondName, &age); err != nil {
			log.Fatal(err)
		}
		employee.DateOfBirth = int64(time.Now().Year()) - age
		employees = append(employees, employee)
	}
	return employees
}

func GetEmployeesByAgeAndSex(db *sql.DB, age int, sex string) []model.EmployeeResponse {
	rows, err := db.Query("SELECT employees.first_name, employees.last_name, employees.second_name, " +
		"employees.age FROM employees WHERE employees.age = $1 AND employees.sex = $2", age, sex)
	if err != nil {
		log.Fatal(err)
	}

	var employees []model.EmployeeResponse
	for rows.Next() {
		var employee model.EmployeeResponse
		var age int64
		if err := rows.Scan(&employee.FirstName, &employee.LastName, &employee.SecondName, &age); err != nil {
			log.Fatal(err)
		}
		employee.DateOfBirth = int64(time.Now().Year()) - age
		employees = append(employees, employee)
	}
	return employees
}