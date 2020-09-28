package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_DeleteMovementOfEmployee(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "DELETE FROM movement_of_employees WHERE employee_number = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("E").WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteMovementOfEmployee(db, "E")
	assert.NoError(t, err)
}

func TestService_DeleteMovementOfInventory(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "DELETE FROM movement_of_inventory WHERE inventory_number = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("E").WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteMovementOfInventory(db, "E")
	assert.NoError(t, err)
}
