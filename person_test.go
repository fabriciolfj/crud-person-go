package main

import (
	"bytes"
	controller "github.com/person/entrypoint"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddPerson(t *testing.T) {
	requestBody := []byte(`{"name":"John Doe", "age": 39}`)

	mux := http.NewServeMux()
	mux.HandleFunc("/person", controller.HandlePeople)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Post(ts.URL+"/person", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
