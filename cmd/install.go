/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Nykenik24/axle/internal/config"
	"github.com/spf13/cobra"
)

var (
	installedTo      []string
	installVerbosely bool
	installDisabled  bool
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs dependencies from config.yaml",
	Long:  `The install command is used to install the dependencies specified in the .axle.yaml file. This command reads the configuration file to identify the packages and dependencies required by the project, and it proceeds to fetch and install them. It ensures that all necessary components are present and up-to-date, allowing the project to work smoothly. This is typically the first step when setting up a new project or when you need to ensure that all required dependencies are properly installed before running or building the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.ReadConfigFrom(".axle.yaml")

		logIfVerbose := func(message string, format ...any) {
			if installVerbosely {
				fmt.Printf(message+"\n", format...)
			}
		}

		if len(installedTo) == 0 {
			for key := range cfg.GetPackageManagers() {
				installedTo = append(installedTo, key)
			}
		}

		for _, managerName := range installedTo {
			if _, err := cfg.GetPackageManager(managerName); err != nil {
				logIfVerbose("--> Manager %s doesn't exist in config", managerName)
				continue
			}

			manager, _ := cfg.GetPackageManager(managerName)

			var baseCommand string
			if _, hasBaseCommand := manager.Commands["base"]; hasBaseCommand {
				baseCommand = manager.Commands["base"]
			} else {
				baseCommand = managerName
			}

			if _, err := exec.LookPath(baseCommand); err != nil {
				logIfVerbose("--> Manager %s was not found in PATH: %v", managerName, err)
				logIfVerbose("\t-> Make sure you have it installed")
				continue
			}

			for pkgName, pkg := range manager.Packages {
				if !pkg.Enabled && !installDisabled {
					continue
				}

				var installCommand string
				if _, hasInstallCommand := manager.Commands["install"]; hasInstallCommand {
					installCommand = manager.Commands["install"]
				} else {
					installCommand = "install {.name}"
				}

				if _, hasInstallWithVersion := manager.Commands["install_with_version"]; hasInstallWithVersion && pkg.Version != "latest" {
					installCommand = manager.Commands["install_with_version"]
				}

				installCommand = strings.ReplaceAll(installCommand, "{.version}", pkg.Version)
				installCommand = strings.ReplaceAll(installCommand, "{.name}", pkgName)

				pathToManager, _ := exec.LookPath(baseCommand)
				command := exec.Command(pathToManager, strings.Split(installCommand, " ")...)
				command.Env = os.Environ()

				logIfVerbose("--> Running command \"%s\"", command.String())

				// All (or almost all) package managers exit with errors when the package doesn't exist
				if out, err := command.CombinedOutput(); err != nil {
					fmt.Printf("--> Could not install package %s with %s\n", pkgName, managerName)
					fmt.Printf("\t-> Detailed error: %v\n", err)
					fmt.Printf("\t-> Output: %s\n", out)
				} else {
					fmt.Printf("--> Installed package %s/%s\n", managerName, pkgName)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringArrayVarP(&installedTo, "managers", "m", []string{}, "Only install in the specified managers")
	installCmd.Flags().BoolVarP(&installVerbosely, "verbose", "V", false, "Log every step and command run")
	installCmd.Flags().BoolVarP(&installDisabled, "include-disabled", "D", false, "Also install disabled packages")
}
