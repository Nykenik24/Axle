/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project by creating a .axle.yaml configuration file",
	Long: `Creates the .axle.yaml file, that stores axle configuration for the project. 
There, package managers and their dependencies, general configurations and more are managed.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteDefaultConfig()
		fmt.Println("Created .axle.yaml at root")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
