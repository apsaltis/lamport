package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

// Config stores configuration options for Lamport
type Config struct {
	Host string
	Port string
}

// ReadConfig returns a Config created from the supplied config file
func ReadConfig(configFile string) (Config, error) {
	var config Config

	_, err := os.Stat(configFile)
	if err != nil {
		return config, errors.New("Config file missing")
	}

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return config, err
	}
	return config, nil
}
