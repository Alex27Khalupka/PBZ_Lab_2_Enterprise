package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/service"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/store"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type APIServer struct {
	config *Config
	router *mux.Router
	Store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.Store = st

	return nil
}

func (s *APIServer) Start() error {

	log.Println("starting API server")

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/divisions", s.handleGetDivisions).Methods(http.MethodGet)
	s.router.Path("/employees").Queries("division_id", "{division_id}").HandlerFunc(s.handleGetEmployeesByDivision).Methods(http.MethodGet)
	s.router.Path("/employees").Queries("age", "{age}", "sex", "{sex}").HandlerFunc(s.handleGetEmployeesByAgeAndSex).Methods(http.MethodGet)
	s.router.HandleFunc("/employees", s.handleGetEmployees).Methods(http.MethodGet)
	s.router.Path("/inventory").Queries("years", "{years}", "division_id", "{division_id}", "name", "{name}").HandlerFunc(s.handleInventoryByYear).Methods(http.MethodGet)
	s.router.HandleFunc("/inventory", s.handleGetInventory).Methods(http.MethodGet)
	s.router.HandleFunc("/repairs", s.handleGetRepairs).Methods(http.MethodGet)
	s.router.HandleFunc("/waybills", s.handleGetWaybills).Methods(http.MethodGet)
	s.router.HandleFunc("/movement_of_employees", s.handleGetMovementOfEmployees).Methods(http.MethodGet)
	s.router.HandleFunc("/divisions/max_repairs_amount", s.handleGetDivisionMaxRepairAmount).Methods(http.MethodGet)
	s.router.HandleFunc("/employees/{division_id}", s.handlePostEmployees).Methods(http.MethodPost)
	s.router.HandleFunc("/inventory/{division_id}", s.handlePostInventory).Methods(http.MethodPost)
	s.router.HandleFunc("/repairs", s.handlePostRepair).Methods(http.MethodPost)
	s.router.HandleFunc("/waybills", s.handlePostWaybill).Methods(http.MethodPost)
	s.router.HandleFunc("/inventory/{inventory_id}", s.handlePutInventory).Methods(http.MethodPut)
	s.router.HandleFunc("/employees/{employee_id}", s.handlePutEmployee).Methods(http.MethodPut)
	s.router.HandleFunc("/employees/{employee_id}", s.handleDeleteEmployee).Methods(http.MethodDelete)
	s.router.HandleFunc("/inventory/{inventory_id}", s.handleDeleteInventory).Methods(http.MethodDelete)
}

func (s *APIServer) handleGetDivisions(w http.ResponseWriter, r *http.Request) {

	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	divisions := model.DivisionsList{Divisions: service.GetDivisions(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(divisions)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetEmployees(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	employees := model.EmployeesList{ResponseEmployees: service.GetEmployees(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetInventory(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	inventory := model.InventoryList{Inventory: service.GetInventory(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(inventory)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetRepairs(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	repairs := model.RepairsList{Repairs: service.GetRepairs(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(repairs)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetWaybills(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	waybills := model.WaybillsList{Waybills: service.GetWaybills(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(waybills)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetMovementOfEmployees(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	documents := model.MovementsOfEmployeesList{MovementOfEmployees: service.GetMovementOfEmployees(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(documents)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetMovementOfInventory(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	movementOfInventory := model.MovementOfInventoryList{MovementOfInventory: service.GetMovementOfInventory(s.Store.GetDB())}

	jsonResponse, err := json.Marshal(movementOfInventory)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetEmployeesByDivision(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}
	paramsList, ok := r.URL.Query()["division_id"]
	divisionID := paramsList[0]
	if strings.Contains(divisionID, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'division_id' is missing", http.StatusBadRequest)
		return
	}
	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params", http.StatusBadRequest)
		return
	}

	//log.Println(divisionID)

	employees := model.EmployeeResponseList{EmployeesByDivisionList: service.GetEmployeesByDivision(s.Store.GetDB(), divisionID)}

	jsonResponse, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetEmployeesByAgeAndSex(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}
	paramsList, ok := r.URL.Query()["age"]
	employeeAgeString := paramsList[0]
	if strings.Contains(employeeAgeString, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'age' is missing for age", http.StatusBadRequest)
		return
	}
	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params", http.StatusBadRequest)
		return
	}

	employeeAge, err := strconv.Atoi(employeeAgeString)
	if err != nil {
		http.Error(w, "Can't convert URL Param 'age' to  int", http.StatusBadRequest)
		return
	}

	paramsList, ok = r.URL.Query()["sex"]
	employeeSex := paramsList[0]
	if strings.Contains(employeeSex, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'sex' is missing", http.StatusBadRequest)
		return
	}

	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params for sex", http.StatusBadRequest)
		return
	}

	employees := model.EmployeeResponseList{EmployeesByDivisionList: service.GetEmployeesByAgeAndSex(s.Store.GetDB(), employeeAge, employeeSex)}

	jsonResponse, err := json.Marshal(employees)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleInventoryByYear(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}
	paramsList, ok := r.URL.Query()["years"]
	yearsStr := paramsList[0]
	if strings.Contains(yearsStr, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'years' is missing", http.StatusBadRequest)
		return
	}
	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params", http.StatusBadRequest)
		return
	}

	years, err := strconv.Atoi(yearsStr)
	if err != nil {
		http.Error(w, "Can't convert URL Param 'years' to  int", http.StatusBadRequest)
		return
	}

	paramsList, ok = r.URL.Query()["division_id"]
	divisionID := paramsList[0]
	if strings.Contains(divisionID, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'division_id' is missing", http.StatusBadRequest)
		return
	}

	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params for division_id", http.StatusBadRequest)
		return
	}

	paramsList, ok = r.URL.Query()["name"]
	inventoryName := paramsList[0]
	if strings.Contains(inventoryName, ";") {
		http.Error(w, "param contains ';'", http.StatusBadRequest)
		return
	}

	if !ok || len(paramsList) < 1 {
		http.Error(w, "Url Param 'name' is missing", http.StatusBadRequest)
		return
	}

	if len(paramsList) > 1 {
		http.Error(w, "To many URL Params for name", http.StatusBadRequest)
		return
	}

	inventoryName = strings.Replace(inventoryName, "_", " ", -1)
	inventoryAmount := service.GetInventoryByYears(s.Store.GetDB(), years, divisionID, inventoryName)

	jsonResponse, err := json.Marshal(inventoryAmount)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleGetDivisionMaxRepairAmount(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	divisionsNames := service.GetDivisionMaxRepairsAmount(s.Store.GetDB())
	log.Println(divisionsNames)

	jsonResponse, err := json.Marshal(divisionsNames)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePostInventory(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	divisionID, err := getID(r, "division_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var inventory model.Inventory
	err = decoder.Decode(&inventory)
	log.Println(inventory)

	if err != nil {
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	if inventory.InventoryNumber == "" || inventory.InventoryName == "" || inventory.InventoryModel == "" {
		http.Error(w, "Some fields are empty", http.StatusBadRequest)
		return
	}

	err = service.POSTInventory(s.Store.GetDB(), inventory, divisionID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Inventory with this id already exists", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(inventory)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePostRepair(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var repair model.Repair
	err := decoder.Decode(&repair)
	log.Println(repair)

	if err != nil {
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	if repair.RepairID == "" || repair.InventoryNumber == "" || repair.EmployeeNumber == "" || repair.RepairType == "" ||
		repair.RepairTime == 0 || repair.WaybillNumber == "" || repair.ServiceStartDay.IsZero() {
		http.Error(w, "Some fields are empty", http.StatusBadRequest)
		return
	}

	err = service.POSTRepair(s.Store.GetDB(), repair)
	if err != nil {
		log.Println(err)
		http.Error(w, "Repair with this id already exists", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(repair)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePostWaybill(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var waybill model.Waybill
	err := decoder.Decode(&waybill)

	if err != nil {
		log.Println(err)
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	if waybill.WaybillNumber == "" || waybill.ReceivingDate.IsZero() || waybill.Price == 0 || waybill.DetailName == "" {
		http.Error(w, "Some fields are empty", http.StatusBadRequest)
		return
	}

	err = service.POSTWaybill(s.Store.GetDB(), waybill)
	if err != nil {
		log.Println(err)
		http.Error(w, "Repair with this id already exists", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(waybill)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePostEmployees(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	divisionID, err := getID(r, "division_id")

	decoder := json.NewDecoder(r.Body)
	var employee model.Employee
	err = decoder.Decode(&employee)

	if err != nil {
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	if employee.EmployeeNumber == "" || employee.FirstName == "" || employee.LastName == "" ||
		employee.SecondName == "" || employee.Position == "" || employee.Age == 0 || employee.Sex == "" {
		http.Error(w, "Some fields are empty", http.StatusBadRequest)
		return
	}

	err = service.POSTEmployee(s.Store.GetDB(), employee, divisionID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(employee)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePutInventory(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	inventoryID, err := getID(r, "inventory_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var inventory model.Inventory
	err = decoder.Decode(&inventory)
	log.Println(inventory)

	if err != nil {
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	updatedInventory, err := service.PUTInventory(s.Store.GetDB(), inventory, inventoryID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Some field have wrong format, or id already exists", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(updatedInventory)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handlePutEmployee(w http.ResponseWriter, r *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	employeeID, err := getID(r, "employee_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var employee model.Employee
	err = decoder.Decode(&employee)

	if err != nil {
		http.Error(w, "Wrong request body", http.StatusBadRequest)
		return
	}

	updatedEmployee, err := service.PUTEmployee(s.Store.GetDB(), employee, employeeID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Some field have wrong format, or id already exists", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(updatedEmployee)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(jsonResponse); err != nil {
		log.Fatal(err)
		return
	}
}

func (s *APIServer) handleDeleteEmployee(w http.ResponseWriter, req *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
	}

	employeeID, err := getID(req, "employee_id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteEmployee(s.Store.GetDB(), employeeID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (s *APIServer) handleDeleteInventory(w http.ResponseWriter, req *http.Request) {
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
	}

	inventoryID, err := getID(req, "inventory_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteInventory(s.Store.GetDB(), inventoryID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// func getID returns id of an object from url
func getID(req *http.Request, idName string) (string, error) {
	if strings.Contains(idName, ";") {
		return "", errors.New("param contains ';'")
	}
	vars := mux.Vars(req)
	id := vars[idName]
	return id, nil
}
