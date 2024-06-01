package main

import (
	"log"

	_ "github.com/person/datasource" // Importa o pacote para executar o init()
	"github.com/person/model"
	"github.com/person/service"
)

func main() {
	p := &model.Person{
		Name: "Fabricio",
		Age:  39,
	}

	result := service.Save(p)

	if result != nil {
		log.Fatalf("failed to save person: %v", result)
		return
	}

	log.Println("Person saved successfully")
}
