package apiserver

import (
	"encoding/json"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/model"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/service"
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/store"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct{
	config *Config
	router *mux.Router
	Store *store.Store
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) configureStore() error{
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil{
		return err
	}
	s.Store = st

	return nil
}

func (s *APIServer) Start() error{

	log.Println("starting API server")

	s.configureRouter()

	if err := s.configureStore(); err!=nil{
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/divisions", s.handleGetDivisions).Methods(http.MethodGet)
	s.router.Path("/employees").Queries("division_id", "{division_id}").HandlerFunc(s.handleGetEmployeesByDivision).Methods(http.MethodGet)
	s.router.HandleFunc("/employees", s.handleGetEmployees).Methods(http.MethodGet)
	s.router.HandleFunc("/inventory", s.handleGetInventory).Methods(http.MethodGet)
	s.router.HandleFunc("/repairs", s.handleGetRepairs).Methods(http.MethodGet)
	s.router.HandleFunc("/waybills", s.handleGetWaybills).Methods(http.MethodGet)
	s.router.HandleFunc("/movement_of_employees", s.handleGetMovementOfEmployees).Methods(http.MethodGet)
	s.router.HandleFunc("/movement_of_inventory", s.handleGetMovementOfInventory).Methods(http.MethodGet)
}

func (s *APIServer) handleGetDivisions(w http.ResponseWriter, r *http.Request){

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


func (s *APIServer) handleGetEmployees(w http.ResponseWriter, r *http.Request){
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("asdasdsadsdasdsad")
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

func (s *APIServer) handleGetInventory(w http.ResponseWriter, r *http.Request){
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

func (s *APIServer) handleGetRepairs(w http.ResponseWriter, r *http.Request){
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

func (s *APIServer) handleGetWaybills(w http.ResponseWriter, r *http.Request){
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

func (s *APIServer) handleGetMovementOfEmployees(w http.ResponseWriter, r *http.Request){
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

func (s *APIServer) handleGetMovementOfInventory(w http.ResponseWriter, r *http.Request){
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

func (s *APIServer) handleGetEmployeesByDivision(w http.ResponseWriter, r *http.Request){
	if err := s.Store.Open(); err != nil {
		log.Fatal(err)
		return
	}
	paramsList, ok := r.URL.Query()["division_id"]
	divisionID := paramsList[0]

	if !ok || len(paramsList) < 1{
		http.Error(w, "Url Param 'division_id' is missing", http.StatusBadRequest)
		return
	}
	if len(paramsList) > 1{
		http.Error(w, "To many URL Params", http.StatusBadRequest)
		return
	}

	log.Println(divisionID)

	employees := model.EmployeeByDivisionList{EmployeesByDivisionList: service.GetEmployeesByDivision(s.Store.GetDB(), divisionID)}

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


// func getID returns id of an object from url
func getID(req *http.Request, idName string) string {
	vars := mux.Vars(req)
	id := vars[idName]
	return id
}