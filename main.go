package main

import (
	"context"
	"github.com/hashicorp/consul/api"
	"github.com/magiconair/properties"
	log "github.com/person/config"
	_ "github.com/person/datasource" // Importa o pacote para executar o init()
	controller "github.com/person/entrypoint"
	_ "github.com/person/redis"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"os/signal"
	"time"
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
		log.Log.Fatal("fail to create consul client", err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "person-service-" + uuid.NewV4().String()
	registration.Name = "person-service"
	registration.Port = 8000
	registration.Address = "localhost"
	registration.Tags = []string{"urlprefix-/person strip=/person"}

	err = consul.Agent().ServiceRegister(registration)

	if err != nil {
		log.Log.Fatal("fail to register service: %v", err)
	}
}

func main() {
	registerService()
	mux := http.NewServeMux()
	http.HandleFunc("/person", controller.RecoveryMiddleware(controller.HandlePeople))
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(":8100", nil); err != nil {
			log.Log.Fatal("fail star server", err)
		}
	}()

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Log.Fatal("Server forced to shutdown: ", err)
	}

	log.Log.Info("Server exiting")
}
