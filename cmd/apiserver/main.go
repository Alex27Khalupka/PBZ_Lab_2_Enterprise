package main

import (
	"github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/apiserver"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	config := apiserver.NewConfig()

	s := apiserver.New(config)

	// starting server
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
//здравствуйте я хотел бы с вами связаться @fanatLYAMYR мой телеграмм ответьсч пожалуйста )

