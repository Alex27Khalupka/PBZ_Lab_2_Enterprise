package service

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestService_UpdateInventory(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "UPDATE inventory SET inventory_number = \\$1, inventory_name = \\$2, inventory_model = \\$3, " +
		"year_of_issue = \\$4 WHERE inventory_number = \\$5"

	date, err := time.Parse(shortForm, "2001-09-18")

	if err != nil {
		log.Fatal(err)
	}

	inventory := model.Inventory{
		InventoryNumber: "I",
		InventoryName:   "name",
		InventoryModel:  "model",
		YearOfIssue:     date,
	}

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("I", "name", "model", date, "I2").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = UpdateInventory(db, inventory, "I2")
	assert.NoError(t, err)
}

func TestService_UpdateEmployee(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "UPDATE employees SET employee_number = \\$1, first_name = \\$2, last_name = \\$3, second_name = \\$4, " +
		"position = \\$5, age = \\$6, sex = \\$7 WHERE employee_number = \\$8"

	employee := model.Employee{
		EmployeeNumber: "1",
		FirstName:      "2",
		LastName:       "3",
		SecondName:     "4",
		Position:       "5",
		Age:            6,
		Sex:            "7",
	}

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("1", "2", "3", "4", "5", 6, "7").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := UpdateEmployee(db, employee, "E0")
	assert.NoError(t, err)
}
