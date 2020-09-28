package service

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestService_GetDivisions(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations.
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	defer db.Close()

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"division_number", "division_name"}).
		AddRow("D4", "Division 4").
		AddRow("D5", "Division 5")

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

func TestService_GetEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_number", "first_name", "last_name", "second_name", "position", "age", "sex"}).
		AddRow("E1", "Alex", "Employee", "First", "leading engineer", 40, "Male").
		AddRow("E2", "Maggie", "Employee", "Second", "engineer", 35, "Female")

	mock.ExpectQuery("^SELECT (.+) FROM employees").WillReturnRows(rows)

	employees := GetEmployees(db)

	expectedEmployees := []model.Employee{
		{
			EmployeeNumber: "E1",
			FirstName:      "Alex",
			LastName:       "Employee",
			SecondName:     "First",
			Position:       "leading engineer",
			Age:            40,
			Sex:            "Male",
		},
		{
			EmployeeNumber: "E2",
			FirstName:      "Maggie",
			LastName:       "Employee",
			SecondName:     "Second",
			Position:       "engineer",
			Age:            35,
			Sex:            "Female",
		},
	}

	assert.Equal(t, expectedEmployees, employees)

}

func TestService_GetInventory(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-17")
	if err != nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-01-07")
	if err != nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"inventory_number", "inventory_name", "inventory_model", "year_of_issue"}).
		AddRow("I1", "Awesome machine", "M111", date1).
		AddRow("I2", "Cool machine", "M222", date2)

	mock.ExpectQuery("^SELECT (.+) FROM inventory").WillReturnRows(rows)

	inventory := GetInventory(db)

	expectedInventory := []model.Inventory{
		{
			InventoryNumber: "I1",
			InventoryName:   "Awesome machine",
			InventoryModel:  "M111",
			YearOfIssue:     date1,
		},
		{
			InventoryNumber: "I2",
			InventoryName:   "Cool machine",
			InventoryModel:  "M222",
			YearOfIssue:     date2,
		},
	}

	assert.Equal(t, expectedInventory, inventory)
}

func TestService_GetRepairs(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err != nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err != nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"repair_id", "inventory_number", "service_start_day", "repair_type", "repair_time",
		"employee_number", "waybill_number"}).
		AddRow("R1", "I1", date1, "minor fix", 1, "E44", "W1").
		AddRow("R2", "I2", date2, "major fix", 2, "E77", "W2")

	mock.ExpectQuery("^SELECT (.+) FROM repairs").WillReturnRows(rows)

	repairs := GetRepairs(db)

	expectedRepairs := []model.Repair{
		{
			RepairID:        "R1",
			InventoryNumber: "I1",
			ServiceStartDay: date1,
			RepairType:      "minor fix",
			RepairTime:      1,
			EmployeeNumber:  "E44",
			WaybillNumber:   "W1",
		},
		{
			RepairID:        "R2",
			InventoryNumber: "I2",
			ServiceStartDay: date2,
			RepairType:      "major fix",
			RepairTime:      2,
			EmployeeNumber:  "E77",
			WaybillNumber:   "W2",
		},
	}

	assert.Equal(t, expectedRepairs, repairs)

}

func TestService_GetWaybills(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err != nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err != nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"waybill_number", "receiving_date", "price", "detail_name"}).
		AddRow("W1", date1, 100, "detail 1").
		AddRow("W2", date2, 200, "detail 2")

	mock.ExpectQuery("^SELECT (.+) FROM waybills").WillReturnRows(rows)

	waybills := GetWaybills(db)

	expectedWaybills := []model.Waybill{
		{
			WaybillNumber: "W1",
			ReceivingDate: date1,
			Price:         100,
			DetailName:    "detail 1",
		},
		{
			WaybillNumber: "W2",
			ReceivingDate: date2,
			Price:         200,
			DetailName:    "detail 2",
		},
	}

	assert.Equal(t, expectedWaybills, waybills)

}

func TestService_GetMovementOfEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err != nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err != nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"employee_number", "movement_date", "division_number"}).
		AddRow("E77", date1, "D2").
		AddRow("E16", date2, "D1")

	mock.ExpectQuery("^SELECT (.+) FROM movement_of_employees").WillReturnRows(rows)

	movementOfEmployees := GetMovementOfEmployees(db)

	expectedMovementOfEmployees := []model.MovementOfEmployees{
		{
			EmployeeNumber: "E77",
			MovementDate:   date1,
			DivisionNumber: "D2",
		},
		{
			EmployeeNumber: "E16",
			MovementDate:   date2,
			DivisionNumber: "D1",
		},
	}

	assert.Equal(t, expectedMovementOfEmployees, movementOfEmployees)
}

func TestService_GetMovementOfInventory(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	date1, err := time.Parse(shortForm, "2020-09-19")
	if err != nil {
		log.Fatal(err)
	}

	date2, err := time.Parse(shortForm, "2020-09-18")
	if err != nil {
		log.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"inventory_number", "movement_date", "division_number"}).
		AddRow("I1", date1, "D1").
		AddRow("I2", date2, "D2")

	mock.ExpectQuery("^SELECT (.+) FROM movement_of_inventory").WillReturnRows(rows)

	movementOfInventory := GetMovementOfInventory(db)

	expectedMovementOfInventory := []model.MovementOfInventory{
		{
			InventoryNumber: "I1",
			MovementDate:    date1,
			DivisionNumber:  "D1",
		},
		{
			InventoryNumber: "I2",
			MovementDate:    date2,
			DivisionNumber:  "D2",
		},
	}

	assert.Equal(t, expectedMovementOfInventory, movementOfInventory)

}

func TestService_GetEmployeesByDivision(t *testing.T) {
	db, mock := NewMock()

	defer db.Close()

	query := "SELECT DISTINCT employees.first_name, employees.last_name, employees.second_name, employees.age " +
		"FROM employees " +
		"INNER JOIN movement_of_employees ON employees.employee_number = movement_of_employees.employee_number " +
		"WHERE division_number = \\(SELECT DISTINCT division_number FROM movement_of_employees " +
		"WHERE movement_of_employees.employee_number = employees.employee_number AND movement_date = " +
		"\\(SELECT MAX\\(movement_date\\) FROM movement_of_employees WHERE employee_number = employees.employee_number\\)\\) " +
		"AND division_number = \\$1"

	rows := sqlmock.NewRows([]string{"first_name", "last_name", "second_name", "age"}).
		AddRow("Kimi", "Raikonnen", "Matias", 99).
		AddRow("Alex", "Khalupka", "Andreevich", 19).
		AddRow("Alex", "Lapitsky", "Evgenevich", 20)

	mock.ExpectQuery(query).WithArgs("D1").WillReturnRows(rows)

	employees := GetEmployeesByDivision(db, "D1")

	assert.NotNil(t, employees)
}

func TestService_GetEmployeesByAgeAndSex(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"first_name", "last_name", "second_name", "age"}).
		AddRow("Alex", "Employee", "First", 40).
		AddRow("Alex", "Employee", "Second", 40)

	mock.ExpectQuery("^SELECT (.+) FROM employees WHERE*").
		WithArgs(40, "Male").
		WillReturnRows(rows)

	employees := GetEmployeesByAgeAndSex(db, 40, "Male")

	expectedEmployees := []model.EmployeeResponse{
		{
			FirstName:   "Alex",
			LastName:    "Employee",
			SecondName:  "First",
			DateOfBirth: 1980,
		},
		{
			FirstName:   "Alex",
			LastName:    "Employee",
			SecondName:  "Second",
			DateOfBirth: 1980,
		},
	}

	assert.Equal(t, expectedEmployees, employees)
}

func TestService_EmployeeByID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_number", "first_name", "last_name", "second_name", "position", "age", "sex"}).
		AddRow("E1", "Alex", "Employee", "First", "leading engineer", 40, "Male")

	mock.ExpectQuery("SELECT employees.employee_number, employees.first_name, employees.last_name, " +
		"employees.second_name, employees.position, employees.age, employees.sex FROM employees " +
		"WHERE employee_number = \\$1").
		WithArgs("E1").
		WillReturnRows(rows)

	expectedEmployee := model.Employee{
		EmployeeNumber: "E1",
		FirstName:      "Alex",
		LastName:       "Employee",
		SecondName:     "First",
		Position:       "leading engineer",
		Age:            40,
		Sex:            "Male",
	}

	employee := GetEmployeeByID(db, "E1")

	assert.Equal(t, expectedEmployee, employee)
}

func TestService_GetMaxRepairs(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"0"})

	mock.ExpectQuery("SELECT MAX\\(count\\) FROM \\(SELECT division_number, count\\(\\*\\) " +
		"FROM division_repair " +
		"GROUP BY division_number\\) AS foo").
		WillReturnRows(rows)

	max := GetMaxRepairs(db)

	expectedMax := int64(0)

	assert.Equal(t, expectedMax, max)
}

func TestService_GetDivisionNumberWithMaxRepairsAmount(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"division_number"}).
		AddRow("D100").
		AddRow("D101")

	mock.ExpectQuery("SELECT division_number FROM \\(SELECT division_number, count\\(\\*\\) FROM " +
		"division_repair " +
		"GROUP BY division_number\\) AS foo WHERE count = \\$1").
		WithArgs(10).
		WillReturnRows(rows)

	numbers := GetDivisionNumberWithMaxRepairsAmount(db, 10)

	expectedNumbers := []string{"D100", "D101"}

	assert.Equal(t, expectedNumbers, numbers)
}

func TestService_GetDivisionNameByID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"division_name"}).
		AddRow("name 1").
		AddRow("name 2")

	mock.ExpectQuery("SELECT division_name FROM divisions WHERE division_number IN \\(\\$1\\)").
		WithArgs("D100").
		WillReturnRows(rows)

	names := GetDivisionNameByID(db, []string{"D100"})

	expectedNames := []string{"name 1", "name 2"}

	assert.Equal(t, expectedNames, names)
}

func TestService_GetInventoryByYear(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"0"})

	years := 3
	currentYear, monthInt, dayInt := time.Now().Date()
	year := strconv.Itoa(currentYear - years)
	month := strconv.Itoa(int(monthInt))
	day := strconv.Itoa(dayInt)
	if len(day) < 2 {
		day = "0" + day
	}

	if len(month) < 2 {
		month = "0" + month
	}

	dateToParse := year + "-" + month + "-" + day
	date, err := time.Parse(shortForm, dateToParse)
	if err != nil {
		log.Fatal(err)
	}

	mock.ExpectQuery("SELECT count\\(\\*\\) FROM movement_of_inventory INNER JOIN inventory ON "+
		"inventory.inventory_number = movement_of_inventory.inventory_number WHERE movement_date \\> \\$1 "+
		"AND inventory_name = \\$2 AND division_number = \\$3").
		WithArgs(date, "name", "D").
		WillReturnRows(rows)

	log.Println(date.Year())
	count := GetInventoryByYears(db, years, "D", "name")

	assert.Equal(t, 0, count)
}
