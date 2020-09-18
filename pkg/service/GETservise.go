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
	rows, err := db.Query("SELECT employees.employee_number, employees.first_name, employees.last_name, employees.second_name, " +
		"employees.position, employees.age, employees.sex FROM employees")
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

func GetInventory(db *sql.DB) []model.Inventory {
	rows, err := db.Query("SELECT inventory.inventory_number, inventory_name, inventory.inventory_model, inventory.year_of_issue FROM inventory")
	if err != nil {
		log.Fatal(err)
	}

	var inventoryList []model.Inventory
	for rows.Next() {
		var inventory model.Inventory
		if err := rows.Scan(&inventory.InventoryNumber, &inventory.InventoryName, &inventory.InventoryModel, &inventory.YearOfIssue); err != nil {
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
		if err := rows.Scan(&waybill.WaybillNumber, &waybill.ReceivingDate, &waybill.Price, &waybill.DetailName); err != nil {
			log.Fatal(err)
		}
		waybillsList = append(waybillsList, waybill)
	}
	return waybillsList
}