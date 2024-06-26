package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Uuid string `json:"uuid"`
}

func getPerson(uuid string) Person {
	resp, err := http.Get("http://localhost:8000/person?uuid=" + uuid)
	if err != nil {
		fmt.Errorf("fail get person uuiid %s, details %s", uuid, err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("erro read body request, details %s", err.Error())
	}

	var person Person
	if err = json.Unmarshal(body, &person); err != nil {
		log.Fatalf("fail deserialize body request, details %s", err.Error())
	}

	return person
}

func main() {
	person := getPerson("0e964249-5681-445e-a916-4357288730fa")
	fmt.Println(person)
}
