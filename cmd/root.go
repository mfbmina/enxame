package cmd

import (
	"fmt"

	"github.com/mfbmina/enxame/reporter"
	"github.com/mfbmina/enxame/swarm"
	"github.com/spf13/cobra"
)

var format string
var requestsPerUser int
var users int

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

	runCmd := runCmd()
	runCmd.Flags().IntVarP(&requestsPerUser, "requests_per_user", "r", 5, "Max number of requests per user")
	runCmd.Flags().IntVarP(&users, "users", "u", 10, "Max number of concurrent users")
	runCmd.Flags().StringVarP(&format, "format", "f", "txt", "The report format (txt, csv, json)")
	rootCmd.AddCommand(runCmd)
}

func runCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run a test",
		Long:  "Run a test against the given URI",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Swarming", args[0], "...")
			responses := swarm.Swarm(args[0], requestsPerUser, users)

			r := reporter.NewReporter(format, "test", responses)
			r.Report()

			fmt.Println("Done!")
		},
	}
}
