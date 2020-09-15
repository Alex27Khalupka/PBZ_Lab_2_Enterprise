package service

import (
	"database/sql"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"log"
)

func GetDivisions(db *sql.DB) []model.Division {

	rows, err := db.Query("SELECT division_number, division_name FROM divisions")
	if err != nil {
		log.Fatal(err)
	}

	var divisions []model.Division
	for rows.Next() {
		var  division model.Division
		if err := rows.Scan(&division.DivisionNumber, &division.DivisionName); err != nil {
			log.Fatal(err)
		}
		divisions = append(divisions, division)
	}

	return divisions
}
