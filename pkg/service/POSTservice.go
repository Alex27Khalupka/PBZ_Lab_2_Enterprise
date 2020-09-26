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

func POSTInventory(db *sql.DB, inventory model.Inventory, divisionID string) error{

	query := "INSERT INTO inventory (inventory_number, inventory_name, inventory_model, year_of_issue) " +
		"VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventory.InventoryNumber, inventory.InventoryName, inventory.InventoryModel,
		inventory.YearOfIssue)

	if err !=nil{
		return err
	}

	if divisionID != "" {
		err = POSTMovementOfInventory(db, inventory.InventoryNumber, divisionID)
	}

	return nil
}

func POSTMovementOfInventory(db *sql.DB, inventoryNumber string, divisionID string) error{
	query := "INSERT INTO movement_of_inventory (inventory_number, movement_date, division_number) VALUES" +
		" ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(inventoryNumber, time.Now().Format("2020-01-01"), divisionID)

	if err !=nil{
		return err
	}

	return nil
}

func POSTRepair(db *sql.DB, repair model.Repair) error{

	query := "INSERT INTO inventory (repairs.repair_id, repairs.inventory_number, repairs.service_start_date, " +
		"repairs.repair_type, repairs.days_to_repair, repairs.employee_number, repairs.waybill_number) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(repair.RepairID, repair.InventoryNumber, repair.ServiceStartDay, repair.RepairType,
		repair.RepairTime, repair.EmployeeNumber, repair.WaybillNumber)

	if err !=nil{
		return err
	}

	return nil
}
