package main

import (
	"github.com/hashicorp/consul/api"
	"github.com/magiconair/properties"
	_ "github.com/person/datasource" // Importa o pacote para executar o init()
	controller "github.com/person/entrypoint"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func registerService() {
	p := properties.MustLoadFile("config.properties", properties.UTF8)

	host := os.Getenv("CONSUL_URL")
	if host == "" {
		host = p.MustGetString("consul_url")
	}

	config := api.DefaultConfig()
	config.Address = host
	consul, err := api.NewClient(config)

	if err != nil {
		logrus.Fatal("fail to create consul client", err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "person-service-" + uuid.NewV4().String()
	registration.Name = "person-service"
	registration.Port = 8000
	registration.Address = "localhost"
	registration.Tags = []string{"urlprefix-/person strip=/person"}

	err = consul.Agent().ServiceRegister(registration)

	if err != nil {
		logrus.Fatal("fail to register service: %v", err)
	}
}

func main() {
	registerService()
	http.HandleFunc("/person", controller.HandlePeople)
	http.Handle("/metrics", promhttp.Handler())

	logrus.Println("Server starting on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		logrus.Fatal("fail star server", err)
	}
}
