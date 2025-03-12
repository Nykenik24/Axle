package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadRoot(path string) ConfigRoot {
	var root ConfigRoot

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &root)
	if err != nil {
		log.Fatalf("Error unmarshalling config YAML: %v", err)
	}

	return root
}

func ReadConfig(root ConfigRoot) Config {
	return Config{
		root: root,
	}
}

func ReadConfigFrom(path string) Config {
	return Config{
		root: ReadRoot(path),
	}
}
