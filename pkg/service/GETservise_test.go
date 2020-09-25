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

func TestService_GetRepairs(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err!=nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err!=nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"repair_id", "inventory_number", "service_start_day", "repair_type", "repair_time",
		"employee_number", "waybill_number"}).
		AddRow( "R1", "I1", date1, "minor fix", 1, "E44", "W1").
		AddRow( "R2", "I2", date2, "major fix", 2, "E77", "W2" )

	mock.ExpectQuery("^SELECT (.+) FROM repairs").WillReturnRows(rows)

	repairs := GetRepairs(db)


	expectedRepairs := []model.Repair{
		{
			RepairID: "R1",
			InventoryNumber: "I1",
			ServiceStartDay: date1,
			RepairType: "minor fix",
			RepairTime: 1,
			EmployeeNumber: "E44",
			WaybillNumber: "W1",
		},
		{
			RepairID: "R2",
			InventoryNumber: "I2",
			ServiceStartDay: date2,
			RepairType: "major fix",
			RepairTime: 2,
			EmployeeNumber: "E77",
			WaybillNumber: "W2",
		},
	}

	assert.Equal(t, expectedRepairs, repairs)

}

func TestService_GetWaybills(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err!=nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err!=nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"waybill_number", "receiving_date", "price", "detail_name"}).
		AddRow( "W1", date1, 100, "detail 1").
		AddRow( "W2", date2, 200, "detail 2")

	mock.ExpectQuery("^SELECT (.+) FROM waybills").WillReturnRows(rows)

	waybills := GetWaybills(db)


	expectedWaybills := []model.Waybill{
		{
			WaybillNumber: "W1",
			ReceivingDate: date1,
			Price: 100,
			DetailName: "detail 1",
		},
		{
			WaybillNumber: "W2",
			ReceivingDate: date2,
			Price: 200,
			DetailName: "detail 2",
		},
	}

	assert.Equal(t, expectedWaybills, waybills)

}

func TestService_GetMovementOfEmployees(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err!=nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err!=nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"employee_number", "movement_date", "division_number"}).
		AddRow( "E77", date1, "D2").
		AddRow( "E16", date2, "D1")

	mock.ExpectQuery("^SELECT (.+) FROM movement_of_employees").WillReturnRows(rows)

	movementOfEmployees := GetMovementOfEmployees(db)


	expectedMovementOfEmployees := []model.MovementOfEmployees{
		{
			EmployeeNumber: "E77",
			MovementDate: date1,
			DivisionNumber: "D2",
		},
		{
			EmployeeNumber: "E16",
			MovementDate: date2,
			DivisionNumber: "D1",
		},
	}

	assert.Equal(t, expectedMovementOfEmployees, movementOfEmployees)
}

func TestService_GetMovementOfInventory(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err!=nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err!=nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"inventory_number", "movement_date", "division_number"}).
		AddRow( "I1", date1, "D1").
		AddRow( "I2", date2, "D2")

	mock.ExpectQuery("^SELECT (.+) FROM movement_of_inventory").WillReturnRows(rows)

	movementOfInventory := GetMovementOfInventory(db)


	expectedMovementOfInventory := []model.MovementOfInventory{
		{
			InventoryNumber: "I1",
			MovementDate: date1,
			DivisionNumber: "D1",
		},
		{
			InventoryNumber: "I2",
			MovementDate: date2,
			DivisionNumber: "D2",
		},
	}

	assert.Equal(t, expectedMovementOfInventory, movementOfInventory)

}

func TestService_GetEmployeesByDivision(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"first_name", "last_name", "second_name", "age"}).
		AddRow( "Kimi", "Raikonnen", "Matias", 99).
		AddRow( "Alex", "Khalupka", "Andreevich", 19).
		AddRow( "Alex", "Lapitsky", "Evgenevich", 20)

	mock.ExpectQuery("^SELECT (.+) FROM employees INNER JOIN movement_of_employees ON " +
		"employees.employee_number = movement_of_employees.employee_number " +
		"WHERE movement_of_employees.division_number = (.+)").WillReturnRows(rows)

	expectedEmployees := []model.EmployeeResponse{
		{
			FirstName: "Kimi",
			LastName: "Raikonnen",
			SecondName: "Matias",
			DateOfBirth: 1921,
		},
		{
			FirstName: "Alex",
			LastName: "Khalupka",
			SecondName: "Andreevich",
			DateOfBirth: 2001,
		},
		{
			FirstName: "Alex",
			LastName: "Lapitsky",
			SecondName: "Evgenevich",
			DateOfBirth: 2000,
		},
	}

	employees := GetEmployeesByDivision(db, "D2")

	assert.Equal(t, expectedEmployees, employees)
}

func TestService_GetEmployeesByAgeAndSex(t *testing.T){
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"first_name", "last_name", "second_name", "age"}).
		AddRow( "Alex", "Employee", "First", 40).
		AddRow( "Alex", "Employee", "Second", 40)

	mock.ExpectQuery("^SELECT (.+) FROM employees WHERE*").
		WithArgs(40, "Male").
		WillReturnRows(rows)

	employees := GetEmployeesByAgeAndSex(db, 40, "Male")

	expectedEmployees := []model.EmployeeResponse{
		{
			FirstName: "Alex",
			LastName: "Employee",
			SecondName: "First",
			DateOfBirth: 1980,
		},
		{
			FirstName: "Alex",
			LastName: "Employee",
			SecondName: "Second",
			DateOfBirth: 1980,
		},
	}

	assert.Equal(t, expectedEmployees, employees)
}
