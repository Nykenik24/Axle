package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	root ConfigRoot
}

func (c Config) GetRoot() ConfigRoot {
	return c.root
}

func (c Config) Write(path string) {
	config := c.GetRoot()

	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshaling config to YAML: %v", err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating config %s: %v", path, err)
	}
	defer file.Close()

	_, err = file.Write([]byte(yamlData))
	if err != nil {
		log.Fatalf("Error writing to config %s: %v", path, err)
	}
}
