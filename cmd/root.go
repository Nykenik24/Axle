package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "axle",
	Short: "Axle - A universal dependency manager",
	Long:  "Axle aggregates multiple package managers (npm, pip, gem, etc.) into a single CLI tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `axle --help` to see available commands.")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
