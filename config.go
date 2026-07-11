package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Target struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Path   string `yaml:"path"`
}

type Route struct {
	Target Target `yaml:"target"`
}

type Config struct {
	Routes map[string]Route `yaml:"routes"`
}

func loadConfig(file_path string) (Config, error) {
	var cfg Config

	file, err := os.Open(file_path)
	if err != nil {
		return cfg, fmt.Errorf("Error decoding YAML: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("Error decoding YAML: %v", err)
	}

	return cfg, nil
}
