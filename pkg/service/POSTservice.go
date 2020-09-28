package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"log"
	"strconv"
	"time"
)

func POSTEmployee(db *sql.DB, employee model.Employee, divisionID string) error {

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

	if err != nil {
		return err
	}

	err = POSTMovementOfEmployees(db, employee.EmployeeNumber, divisionID)

	return nil
}

func POSTMovementOfEmployees(db *sql.DB, employeeNumber string, divisionID string) error {
	query := "INSERT INTO movement_of_employees (employee_number, movement_date, division_number) VALUES" +
		" ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(employeeNumber, time.Now().Format("2006-01-02"), divisionID)

	if err != nil {
		return err
	}

	return nil
}

func POSTInventory(db *sql.DB, inventory model.Inventory, divisionID string) error {

	query := "INSERT INTO inventory (inventory_number, inventory_name, inventory_model, year_of_issue) " +
		"VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventory.InventoryNumber, inventory.InventoryName, inventory.InventoryModel,
		inventory.YearOfIssue)

	if err != nil {
		return err
	}

	if divisionID != "" {
		err = POSTMovementOfInventory(db, inventory.InventoryNumber, divisionID)
	}

	return nil
}

func POSTMovementOfInventory(db *sql.DB, inventoryNumber string, divisionID string) error {
	query := "INSERT INTO movement_of_inventory (inventory_number, movement_date, division_number) VALUES" +
		" ($1, $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(inventoryNumber, time.Now().Format("2006-01-02"), divisionID)

	if err != nil {
		return err
	}

	return nil
}

func POSTRepair(db *sql.DB, repair model.Repair) error {

	if err := CreateRepair(db, repair); err != nil {
		return err
	}

	if err := CreateDivisionRepair(db, repair); err != nil {
		return err
	}

	return nil
}

func CreateDivisionRepair(db *sql.DB, repair model.Repair) error {

	query := "INSERT INTO division_repair (division_number, repair_id) VALUES ($1, $2)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	divisionID, err := GetInventoryDivisionID(db, repair.InventoryNumber)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(divisionID, repair.RepairID)

	if err != nil {
		return err
	}

	return nil
}

func GetInventoryDivisionID(db *sql.DB, inventoryNumber string) (string, error) {
	query := "SELECT DISTINCT movement_of_inventory.division_number FROM inventory " +
		"INNER JOIN movement_of_inventory ON inventory.inventory_number = movement_of_inventory.inventory_number " +
		"WHERE division_number = " +
		"(SELECT DISTINCT division_number FROM movement_of_inventory " +
		"WHERE movement_of_inventory.inventory_number = inventory.inventory_number AND movement_date = " +
		"(SELECT MAX(movement_date) FROM movement_of_inventory WHERE inventory_number = inventory.inventory_number)) " +
		"AND inventory.inventory_number = $1"

	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := db.Query(query, inventoryNumber)

	if err != nil {
		return "", err
	}

	var divisionID string
	for rows.Next() {
		err := rows.Scan(&divisionID)
		if err != nil {
			log.Fatal(err)
		}
	}

	//log.Println("kek : ", inventoryNumber)
	return divisionID, nil
}

func CreateRepair(db *sql.DB, repair model.Repair) error {

	query := "INSERT INTO repairs (repair_id, inventory_number, service_start_date, " +
		"repair_type, days_to_repair, employee_number, waybill_number) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	daysToRepair := strconv.Itoa(int(repair.RepairTime)) + " days"
	//log.Println(string(repair.RepairTime))

	_, err = stmt.Exec(repair.RepairID, repair.InventoryNumber, repair.ServiceStartDay, repair.RepairType,
		daysToRepair, repair.EmployeeNumber, repair.WaybillNumber)

	if err != nil {
		return err
	}

	return nil
}

func POSTWaybill(db *sql.DB, waybill model.Waybill) error {

	query := "INSERT INTO waybills (waybill_number, receiving_date, price, detail_name) " +
		"VALUES ($1, $2, $3, $4)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(waybill.WaybillNumber, waybill.ReceivingDate, waybill.Price, waybill.DetailName)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
