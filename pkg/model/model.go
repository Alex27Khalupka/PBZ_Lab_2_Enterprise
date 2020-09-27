package model

import (
	"time"
)

type Inventory struct {
	InventoryNumber string    `json:"inventory_number"`
	InventoryName   string    `json:"inventory_name"`
	InventoryModel  string    `json:"inventory_model"`
	YearOfIssue     time.Time `json:"year_of_issue"`
}

type MovementOfInventory struct {
	InventoryNumber string    `json:"inventory"`
	MovementDate    time.Time `json:"movement_date"`
	DivisionNumber  string    `json:"division_number"`
}

type Division struct {
	DivisionNumber string `json:"division_number"`
	DivisionName   string `json:"division_name"`
}

type Employee struct {
	EmployeeNumber string `json:"employee_number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	SecondName     string `json:"second_name"`
	Position       string `json:"position"`
	Age            int64  `json:"age"`
	Sex            string `json:"sex"`
}

type Repair struct {
	RepairID        string    `json:"repair_id"`
	InventoryNumber string    `json:"inventory_number"`
	ServiceStartDay time.Time `json:"service_start_date"`
	RepairType      string    `json:"repair_type"`
	RepairTime      int64     `json:"days_to_repair"`
	EmployeeNumber  string    `json:"employee_number"`
	WaybillNumber   string    `json:"waybill_number"`
}

type Waybill struct {
	WaybillNumber string    `json:"waybill_number"`
	ReceivingDate time.Time `json:"receiving_date"`
	Price         int64     `json:"price"`
	DetailName    string    `json:"detail_name"`
}

type MovementOfEmployees struct {
	EmployeeNumber string    `json:"employee_number"`
	MovementDate   time.Time `json:"movement_date"`
	DivisionNumber string    `json:"division_number"`
}

type DivisionsList struct {
	Divisions []Division `json:"divisions"`
}

type EmployeesList struct {
	ResponseEmployees []Employee `json:"employees"`
}

type MovementOfInventoryList struct {
	MovementOfInventory []MovementOfInventory `json:"movement_od_inventory"`
}

type RepairsList struct {
	Repairs []Repair `json:"repairs"`
}

type WaybillsList struct {
	Waybills []Waybill `json:"waybills"`
}

type MovementsOfEmployeesList struct {
	MovementOfEmployees []MovementOfEmployees `json:"documents"`
}

type InventoryList struct {
	Inventory []Inventory `json:"inventory"`
}

type EmployeeResponse struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	SecondName     string `json:"second_name"`
	DateOfBirth    int64  `json:"date_of_birth"`
}

type EmployeeResponseList struct{
	EmployeesByDivisionList []EmployeeResponse `json:"employees_by_division"`
}