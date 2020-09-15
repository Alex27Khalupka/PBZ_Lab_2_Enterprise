package store

import (
	"database/sql"
	"log"
	"os"
)

type Store struct{
	config *Config
	db	   *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error{
	dbURI := os.Getenv("POSTGRES_URI")
	log.Println("Database URI:", dbURI)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil{
		return err
	}

	s.db = db

	return nil
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

func (s *Store) Close(){
	_ = s.db.Close()
}


