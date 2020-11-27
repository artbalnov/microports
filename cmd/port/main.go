package main

import (
	"flag"
	"log"

	"github.com/microports/app/service/port"
	"github.com/microports/app/util/env"
)

var (
	envFile string
	address string
)

func welcome() {
	log.Printf("[init] Copyright (C) 2020 Artemy Balnov. All Rights Reserve. Start Port service")
}

func main() {
	// Load flags
	flag.StringVar(&envFile, "env", "", "File with environment variables")
	flag.StringVar(&address, "address", ":8080", "Service address")

	flag.Parse()

	// Abandon all hope ye who enter here
	welcome()

	// Load env file
	if envFile != "" {
		err := env.LoadEnvFile(envFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create service
	service, err := port.Factory()
	if err != nil {
		log.Fatal(err)
	}

	err = service.Attach(address)
	if err != nil {
		log.Fatal(err)
	}
}
