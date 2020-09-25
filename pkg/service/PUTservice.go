package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
)

func PUTInventory(db *sql.DB, inventoryToUpdate model.Inventory, inventoryID string) (model.Inventory, error){

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

	query := "UPDATE inventory SET inventory_number = $1, inventory_name = $2, inventory_model = $3, " +
		"year_of_issue = $4 WHERE inventory_number = $5"

	stmt, err := db.Prepare(query)
	if err != nil {
		return inventory,err
	}
	defer stmt.Close()

	_, err = stmt.Exec(inventoryToUpdate.InventoryNumber, inventoryToUpdate.InventoryName,
		inventoryToUpdate.InventoryModel, inventoryToUpdate.YearOfIssue, inventoryID)

	return inventoryToUpdate, err
}