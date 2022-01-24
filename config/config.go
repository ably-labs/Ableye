package config

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/kelseyhightower/envconfig"
)

// Config is used to keep track of configuration in one place.
type Config struct {
	DebugMode        bool
	Version          string `default:"0.0.1"`
	AblyGoSDKVersion string
	Key              string `envconfig:"ABLY_PRIVATE_KEY"`
}

// Cfg holds the current configuration.
var Cfg Config

// Initialisation of default configuration.
func init() {
	Cfg.DebugMode = false // Turns debug mode on.

	//Process environment variables and store them in the global cfg.
	if err := envconfig.Process("", &Cfg); err != nil {
		log.Fatal(err.Error())
	}

	if Cfg.Key == "" {
		log.Fatal("ABLY_PRIVATE_KEY not found. Environment variable must be set.")
	}

	// Note: a Github issue has been created to add the ability to get SDK version to
	// all client libraries. See https://github.com/ably/docs/issues/1291 while this
	// functionality is not implemented in the SDK, the only place this information is
	// available is from the go.mod file.

	// Execute a bash command to get the version of ably-go currently in use.
	cmd := "go list -m all | grep ably-go | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		Cfg.AblyGoSDKVersion = "unknown"
		return
	}

	Cfg.AblyGoSDKVersion = string(out)
}
