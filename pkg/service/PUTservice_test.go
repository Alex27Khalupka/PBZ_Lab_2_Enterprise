package service

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestService_PUTInventory(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "UPDATE inventory SET inventory_number = \\$1, inventory_name = \\$2, inventory_model = \\$3, " +
		"year_of_issue = \\$4 WHERE inventory_number = \\$5"

	date, err := time.Parse(shortForm, "2001-09-18")

	if err != nil{
		log.Fatal(err)
	}

	inventory := model.Inventory{
		InventoryNumber: "I",
		InventoryName:   "name",
		InventoryModel:  "model",
		YearOfIssue:     date,
	}

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("I", "name", "model", date).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = PUTInventory(db, inventory, "I2")
	assert.NoError(t, err)
}