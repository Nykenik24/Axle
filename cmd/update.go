/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/spf13/cobra"
)

var (
	targetUpdatedManagers []string
	installAfterUpdate    bool
	updateVerbosely       bool
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates dependencies",
	Long:  `The update command is used to update the existing dependencies to their latest compatible versions. It updates all packages that are not locked to an specific version to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")
		packageManagers := cfg.GetPackageManagers()

		if len(targetUpdatedManagers) == 0 {
			for key := range cfg.GetPackageManagers() {
				targetUpdatedManagers = append(targetUpdatedManagers, key)
			}
		}

		for managerName, packageManager := range packageManagers {
			for _, targetManager := range targetUpdatedManagers {
				if managerName == targetManager {
					for pkgName, pkg := range packageManager.Packages {
						if !pkg.Locked {
							if pkg.Version != "latest" {
								pkg.Version = "latest"
								if updateVerbosely {
									fmt.Printf("--> Updated package %s to latest\n", pkgName)
								}
							}
						} else {
							if updateVerbosely {
								fmt.Printf("--> Skipped package %s: it's locked\n", pkgName)
							}
						}
					}
				}
			}
		}

		cfg.Write(".axle.yaml")
		fmt.Println("--> Finished updating")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringArrayVarP(&targetUpdatedManagers, "managers", "m", []string{}, "Managers that will be updated")
	updateCmd.Flags().BoolVarP(&updateVerbosely, "verbose", "V", false, "Show every updated package")
}
