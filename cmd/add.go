/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/spf13/cobra"
)

var (
	addedTo      string
	addVerbosely bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new dependency to the project",
	Long:  `The add command allows you to add a new dependency to the project by specifying the package name and version. This command updates the project's configuration file .axle.yaml with the new dependency, ensuring it is included in the dependency tree for the project. After running this command, the new dependency will be available for installation or use in your project’s environment.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")
		// packageName := args[0]
		// if _, managerExists := cfg.GetRoot().PackageManagers[addedTo]; !managerExists {
		// 	log.Fatalf("Manager %s doesn't exist in configuration", addedTo)
		// }
		// cfg.GetRoot().PackageManagers[addedTo].Packages[packageName] = config.Package{
		// 	Version: version,
		// 	Enabled: !disabledByDefault,
		// }
		// cfg.Write(".axle.yaml")
		//
		// fmt.Printf("Added dependency %s to %s\n", packageName, addedTo)

		var packageNames []string
		packageNames = append(packageNames, args...)

		if _, managerExists := cfg.GetRoot().PackageManagers[addedTo]; !managerExists {
			log.Fatalf("Manager %s doesn't exist in configuration", addedTo)
		}

		for _, packageName := range packageNames {
			// What we are basically doing is having a nice way of specifying our dependency versions:
			// It splits the string by @ and uses the string at the left as name and the string at the right as version.
			options := strings.Split(packageName, "@")
			name := options[0]
			var version string
			if len(options) == 2 {
				version = options[1]
			} else {
				version = "latest"
			}
			if _, alreadyExists := cfg.GetRoot().PackageManagers[addedTo].Packages[name]; alreadyExists {
				fmt.Printf("WARN: Package %s already exists in %s\n", name, addedTo)
			}
			cfg.GetRoot().PackageManagers[addedTo].Packages[name] = config.Package{
				Enabled: true,
				Version: version,
			}

			if addVerbosely {
				fmt.Printf("--> Added dependency %s to %s\n", name, addedTo)
				if version != "latest" {
					fmt.Printf("\t-> Version: %s\n", version)
				}
			}
		}
		cfg.Write(".axle.yaml")
		fmt.Println("Added all dependencies")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addedTo, "manager", "m", "", "The manager to add the package to")
	addCmd.MarkFlagRequired("manager")

	addCmd.Flags().BoolVarP(&addVerbosely, "verbose", "v", false, "Show all dependencies added")
}
