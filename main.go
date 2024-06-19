package main

import (
	"log"
	"net/http"

	_ "github.com/person/datasource" // Importa o pacote para executar o init()
	controller "github.com/person/entrypoint"
)

func main() {
	http.HandleFunc("/person", controller.HandlePeople)

	log.Println("Server starting on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("fail star server", err)
	}
}
