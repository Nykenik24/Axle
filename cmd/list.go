/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/Nykenik24/axle/internal/lists"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists installed packages",
	Long: `The list command displays a list of all the packages that are currently installed in the project. It provides a detailed overview of the project’s dependencies, showing each package’s name, version, and any other relevant information. This command is useful for tracking which packages are part of the project and verifying that the correct versions are installed. It also helps when you need to audit the project’s dependencies or troubleshoot potential issues with missing or incompatible packages.

For a tree view of all dependencies, see "graph"`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")
		packageList := make([]any, len(cfg.GetPackages()))
		for _, packageManager := range cfg.GetPackageManagers() {
			for name, pkg := range packageManager.Packages {
				var versionColor string
				if pkg.Version == "latest" {
					versionColor = "\033[35m"
				} else {
					versionColor = "\033[32m"
				}
				var enabled string
				if pkg.Enabled {
					enabled = "enabled"
				} else {
					enabled = "disabled"
				}
				packageList = append(packageList, fmt.Sprintf("\033[94m%s\033[0m@%s%s\033[0m \033[90m<%v>\033[0m", name, versionColor, pkg.Version, enabled))
			}
		}
		fmt.Println(lists.UnorderedList(packageList, "- "))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
