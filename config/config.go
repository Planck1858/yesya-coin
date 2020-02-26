package config

import ()

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Host              string
	Port              int
	User              string
	Password          string
	DbName            string
}

func GetConfig(filename string) Configuration {
	var configuration Configuration
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
