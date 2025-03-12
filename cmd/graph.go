/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Displays the dependency tree",
	Long:  `The graph command generates a visual representation of the project’s dependency tree. This allows developers to understand how different dependencies are interconnected and how they affect each other. The graph helps to visualize the structure of the project's dependencies, showing which packages depend on others and making it easier to identify potential issues or conflicts. This command is invaluable for debugging and optimizing dependencies, as it provides a hierarchical view of the project’s ecosystem.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graph called")
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
