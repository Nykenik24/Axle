/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dependencies from config.yaml",
	Long:  `The install command is used to install the dependencies specified in the .axle.yaml file. This command reads the configuration file to identify the packages and dependencies required by the project, and it proceeds to fetch and install them. It ensures that all necessary components are present and up-to-date, allowing the project to work smoothly. This is typically the first step when setting up a new project or when you need to ensure that all required dependencies are properly installed before running or building the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
