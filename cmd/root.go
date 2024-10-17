package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "enxame",
	Short: "A http load tester and benchmarking utility made in Go",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(runCmd())
}
