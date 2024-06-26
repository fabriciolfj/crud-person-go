package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	c := cron.New()
	_, err := c.AddFunc("@every 5s", func() {
		fmt.Println("executing")
	})

	if err != nil {
		log.Fatalf("failed to add cron job: %v", err)
	}

	c.Start()

	select {}
}
