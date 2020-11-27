package test

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	var envFile string

	flag.StringVar(&envFile, "env", "../test.env", "File with environment variables")
	flag.Parse()

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}
}
