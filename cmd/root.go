package cmd

import (
	"github.com/mfbmina/enxame/lib"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
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

	// add new commands here
	rootCmd.AddCommand(runCmd())
}

func runCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run a test",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Service(args[0])
		},
	}
}
