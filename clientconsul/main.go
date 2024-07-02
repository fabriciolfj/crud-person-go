package main

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	consul, err := api.NewClient(config)
	if err != nil {
		log.Fatal("consul client error:", err)
	}

	services, _, err := consul.Health().Service("person-service", "", true, nil)
	if err != nil {
		log.Fatal("consul agent error:", err)
	}

	for _, service := range services {
		log.Println("service:", service.Service.Service)
	}

}
