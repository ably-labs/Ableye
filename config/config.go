package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Config is used to keep track of configuration in one place.
type Config struct {
	DebugMode bool
	Version   string `default:"0.0.1"`
	Key       string `envconfig:"ABLY_PRIVATE_KEY"`
}

// Cfg holds the current configuration.
var Cfg Config

// Initialisation of default configuration.
func init() {
	Cfg.DebugMode = true // Turns debug mode on.

	//Process environment variables and store them in the global cfg.
	if err := envconfig.Process("", &Cfg); err != nil {
		log.Fatal(err.Error())
	}

	if Cfg.Key == ""{
		log.Fatal("ABLY_PRIVATE_KEY not found. Environment variable must be set.")
	}
}
