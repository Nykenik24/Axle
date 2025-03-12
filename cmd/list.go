/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists installed packages",
	Long: `The list command displays a list of all the packages that are currently installed in the project. It provides a detailed overview of the project’s dependencies, showing each package’s name, version, and any other relevant information. This command is useful for tracking which packages are part of the project and verifying that the correct versions are installed. It also helps when you need to audit the project’s dependencies or troubleshoot potential issues with missing or incompatible packages.

For a tree view of all dependencies and their dependencies, see "graph"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
