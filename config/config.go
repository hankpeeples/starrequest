package config

import (
	"embed"

	"gopkg.in/yaml.v3"
)

// Embed the config file so we can use it from any working directory
//
//go:embed config.yaml
var configFile embed.FS

// Config represents the config file
type Config struct {
	Env string `yaml:"env"`
	Log struct {
		Level  string `yaml:"level"`
		Output string `yaml:"output"`
	} `yaml:"log"`
}

// LoadConfig loads the config from the config.yaml file
func LoadConfig() (*Config, error) {
	yamlFile, err := configFile.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
