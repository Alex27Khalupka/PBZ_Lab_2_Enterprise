package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/moemoe89/go-unit-test-sql/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestService_POSTEmployee(t *testing.T) {

	db, mock := NewMock()

	defer db.Close()

	employee := model.Employee{
		EmployeeNumber: "NUM",
		FirstName:      "first_name",
		LastName:       "last_name",
		SecondName:     "second_name",
		Position:       "position",
		Age:            20,
		Sex:            "Female",
	}

	query := "INSERT INTO employees \\(employee_number, first_name, last_name, second_name, position, age, sex\\) " +
		"VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.EmployeeNumber, employee.FirstName, employee.LastName, employee.SecondName,
		employee.Position, employee.Age, employee.Sex).WillReturnResult(sqlmock.NewResult(0, 1))

	err := POSTEmployee(db, employee, "D5")
	assert.NoError(t, err)

}

// a failing test case
func TestService_POSTEmployeeFail(t *testing.T) {

	db, mock := NewMock()

	defer db.Close()

	employee := model.Employee{
		EmployeeNumber: "NUM",
		FirstName:      "first_name",
		LastName:       "last_name",
		SecondName:     "second_name",
		Position:       "position",
		Age:            20,
		Sex:            "Female",
	}

	query := "INSERT INTO employees \\(employee_number, first_name, last_name, second_name, position, age, sex\\) " +
		"VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7\\)"


	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.EmployeeNumber, employee.FirstName, employee.LastName, employee.SecondName,
		employee.Position, "not int", employee.Sex).WillReturnResult(sqlmock.NewResult(0, 1))


	err := POSTEmployee(db, employee, "D5")

	assert.Error(t, err)

}


func TestService_POSTMovementOFEmployees(t *testing.T) {

	db, mock := NewMock()

	defer db.Close()

	employee := model.Employee{
		EmployeeNumber: "NUM",
		FirstName:      "first_name",
		LastName:       "last_name",
		SecondName:     "second_name",
		Position:       "position",
		Age:            20,
		Sex:            "Female",
	}

	query := "INSERT INTO movement_of_employees \\(employee_number, movement_date, division_number\\) VALUES " +
	"\\(\\$1, \\$2, \\$3\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.EmployeeNumber, time.Now().Format("2020-01-01"), "D5").
	WillReturnResult(sqlmock.NewResult(0, 1))

	err := POSTMovementOfEmployees(db, employee.EmployeeNumber, "D5")
	assert.NoError(t, err)
}

func TestService_POSTMovementOFEmployeesFail(t *testing.T) {

	db, mock := NewMock()

	defer db.Close()


	query := "INSERT INTO movement_of_employees \\(employee_number, movement_date, division_number\\) VALUES " +
		"\\(\\$1, \\$2, \\$3\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("EMP", "kek", "D5").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := POSTMovementOfEmployees(db, "kek", "D5")
	assert.Error(t, err)
}