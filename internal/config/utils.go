package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func GetDefaultConfig() Config {
	return Config{
		root: ConfigRoot{
			PackageManagers: make(map[string]PackageManager),
			Settings:        Settings{},
			Graph:           GraphSettings{},
		},
	}
}

func WriteDefaultConfig() {
	config := GetDefaultConfig().GetRoot()

	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshaling default config to YAML: %v", err)
	}

	file, err := os.Create(".axle.yaml")
	if err != nil {
		log.Fatalf("Error creating default config: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(yamlData))
	if err != nil {
		log.Fatalf("Error writing to default config: %v", err)
	}
}
