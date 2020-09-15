package apiserver

import (
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

func (s *APIServer) configureRouter(){

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
