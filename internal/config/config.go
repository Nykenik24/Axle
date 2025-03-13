package config

import (
	"fmt"
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

func (c Config) GetPackageManagers() map[string]PackageManager {
	return c.root.PackageManagers
}

func (c Config) GetPackageManager(name string) (PackageManager, error) {
	packageManager, exists := c.GetPackageManagers()[name]
	if !exists {
		return PackageManager{
			Name:     "none",
			Commands: map[string]string{},
			Packages: map[string]Package{},
		}, fmt.Errorf("Package manager %s doesn't exist", name)
	} else {
		return packageManager, nil
	}
}

func (c Config) GetSettings() Settings {
	return c.root.Settings
}

func (c Config) GetPackages() map[string][]Package {
	packageManagers := c.GetPackageManagers()
	packages := map[string][]Package{}
	for managerName, manager := range packageManagers {
		managerPackages := []Package{}
		for name, pkg := range manager.Packages {
			pkg.Name = name
			managerPackages = append(managerPackages, pkg)
		}
		packages[managerName] = managerPackages
	}
	return packages
}

func (c Config) Write(path string) {
	root := c.GetRoot()

	yamlData, err := yaml.Marshal(&root)
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
