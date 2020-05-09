package config

import (
	"encoding/json"
	"os"
)

// Configuration has config from config.{env}.json file
type Configuration struct {
	Mongo string
}

// GetConfiguration returns config from config.{env}.json file
func GetConfiguration() Configuration {
	var configuration Configuration
	file, err := os.Open("./config/config.development.json")

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
