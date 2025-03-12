/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/spf13/cobra"
)

var (
	removedFrom         string
	removeInteractively bool
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a dependency",
	Long:  `The remove command allows you to remove a specified dependency from the project. When executed, it will completely uninstall the package from the environment, cleaning up any files or references associated with it. This command is useful when you no longer need a particular dependency, whether it’s because it’s no longer required or because you're replacing it with an alternative package. It ensures that the project environment stays lean and free of unnecessary packages.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")
		packageName := args[0]
		if _, managerExists := cfg.GetRoot().PackageManagers[removedFrom]; !managerExists {
			log.Fatalf("Manager %s doesn't exist in configuration", removedFrom)
		}
		if _, packageExists := cfg.GetRoot().PackageManagers[removedFrom].Packages[packageName]; !packageExists {
			log.Fatalf("Package %s can't be removed from %s's dependencies: the package doesn't exist", packageName, removedFrom)
		}

		if removeInteractively {
			var sure string

			fmt.Printf("Are you sure you want to remove %s/%s [y/n]? ", removedFrom, packageName)
			fmt.Scanln(&sure)
			switch strings.ToLower(sure) {
			case "y":
				delete(cfg.GetRoot().PackageManagers[removedFrom].Packages, packageName)
				cfg.Write(".axle.yaml")
				fmt.Printf("Removed dependency %s from %s\n", packageName, removedFrom)
				os.Exit(0)
			case "n":
				fmt.Printf("Cancelled removal of %s/%s", removedFrom, packageName)
				os.Exit(0)
			default:
				fmt.Printf("Invalid value: %s (expected \"y\" or \"n\")", strings.ToLower(sure))
				os.Exit(1)
			}

		} else {
			delete(cfg.GetRoot().PackageManagers[removedFrom].Packages, packageName)
			cfg.Write(".axle.yaml")
			fmt.Printf("Removed dependency %s from %s\n", packageName, removedFrom)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&removedFrom, "manager", "m", "", "The manager to remove the package from")
	removeCmd.MarkFlagRequired("manager")

	removeCmd.Flags().BoolVarP(&removeInteractively, "interactive", "i", false, "Prompt for removal before removing the dependency")
}
