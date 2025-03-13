/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/Nykenik24/axle/internal/lists"
	"github.com/spf13/cobra"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Displays the project's direct dependencies in a tree format",
	Long:  `The tree command shows a simple, tree-like view of the project's direct dependencies. It lists the main packages that the project depends on without going into nested dependencies. This command is helpful for quickly understanding what libraries or tools are included in the project and how they’re organized. It's a straightforward way to see the primary dependencies at a glance.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")
		packages := cfg.GetPackages()
		var tree strings.Builder

		for managerName, managerPackages := range packages {
			tree.Write([]byte("- \033[94m" + managerName + "\033[0m\n"))

			var convertedManagerPackages []any

			var versionColor string
			var enabled string
			for _, pkg := range managerPackages {
				if pkg.Version == "latest" {
					versionColor = "\033[35m"
				} else {
					versionColor = "\033[32m"
				}
				if pkg.Enabled {
					enabled = "enabled"
				} else {
					enabled = "disabled"
				}

				convertedManagerPackages = append(convertedManagerPackages, fmt.Sprintf("\033[94m%s\033[0m@%s%s\033[0m \033[90m<%v>\033[0m", pkg.Name, versionColor, pkg.Version, enabled))
			}
			tree.Write([]byte(lists.UnorderedList(convertedManagerPackages, " - ")))
		}

		fmt.Print(tree.String())
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
