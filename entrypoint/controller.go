package controller

import (
	"encoding/json"
	"github.com/person/service"
	"net/http"

	_ "github.com/person/datasource" // Importa o pacote para executar o init()
	"github.com/person/model"
)

func getByUuid(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	person, err := service.Find(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if person.ID == 0 {
		http.Error(w, "person not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := service.Save(&person)
	if result != nil {
		http.Error(w, result.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlePeople(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createPerson(w, r)
	case "GET":
		getByUuid(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
