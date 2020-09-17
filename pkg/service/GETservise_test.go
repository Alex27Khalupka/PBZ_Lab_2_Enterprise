package service

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	shortForm = "2006-01-02"
)

func TestService_GetDivisions(t *testing.T){
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"division_number", "division_name"}).
		AddRow( "D4", "Division 4").
		AddRow( "D5", "Division 5")

	// This is most important part in our test. Here, literally, we are altering SQL query from MenuByNameAndLanguage
	// function and replacing result with our expected result.
	mock.ExpectQuery("^SELECT (.+) FROM divisions").WillReturnRows(rows)


	// Calls MenuByNameAndLanguage with mocked database connection in arguments list.
	divisions := GetDivisions(db)

	// Here we just construction our expecting result.
	expectedDivisions := []model.Division{
			{
				DivisionNumber: "D4",
				DivisionName:   "Division 4",
			},
			{
				DivisionNumber: "D5",
				DivisionName:   "Division 5",
			},
	}

	assert.Equal(t, expectedDivisions, divisions)

}

func TestService_GetEmployees(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_number", "first_name", "last_name", "second_name", "position", "age", "sex"}).
		AddRow( "E1", "Alex", "Employee", "First", "leading engineer", 40, "Male").
		AddRow( "E2", "Maggie", "Employee", "Second", "engineer", 35, "Female")

	mock.ExpectQuery("^SELECT (.+) FROM employees").WillReturnRows(rows)

	employees := GetEmployees(db)

	expectedEmployees := []model.Employee{
		{
			EmployeeNumber: "E1",
			FirstName: "Alex",
			LastName: "Employee",
			SecondName: "First",
			Position: "leading engineer",
			Age: 40,
			Sex: "Male",
		},
		{
			EmployeeNumber: "E2",
			FirstName: "Maggie",
			LastName: "Employee",
			SecondName: "Second",
			Position: "engineer",
			Age: 35,
			Sex: "Female",
		},
	}

	assert.Equal(t, expectedEmployees, employees)

}

func TestService_GetInventory(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-17")
	if err!=nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-01-07")
	if err!=nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"inventory_number", "inventory_name", "inventory_model", "year_of_issue"}).
		AddRow( "I1", "Awesome machine", "M111", date1).
		AddRow( "I2", "Cool machine", "M222", date2)

	mock.ExpectQuery("^SELECT (.+) FROM inventory").WillReturnRows(rows)

	inventory := GetInventory(db)


	expectedInventory := []model.Inventory{
		{
			InventoryNumber: "I1",
			InventoryName:   "Awesome machine",
			InventoryModel:  "M111",
			YearOfIssue: date1,
		},
		{
			InventoryNumber: "I2",
			InventoryName: "Cool machine",
			InventoryModel: "M222",
			YearOfIssue: date2,
		},
	}

	assert.Equal(t, expectedInventory, inventory)

}