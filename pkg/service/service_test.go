package service

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
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
