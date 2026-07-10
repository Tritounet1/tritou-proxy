package main

import (
	"log"
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

func loadConfig() Config {
	var cfg Config

	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Error decoding YAML: %v", err)
	}

	return cfg
}
