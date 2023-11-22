package cmd

import (
	"fmt"

	"github.com/mfbmina/enxame/lib/reporter"
	"github.com/mfbmina/enxame/lib/swarm"
	"github.com/spf13/cobra"
)

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
	rootCmd.AddCommand(runCmd)
}

func runCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run a test",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Swarming", args[0], "...")
			responses := swarm.Swarm(args[0], requestsPerUser, users)

			// I still need to figure out how to instanciate the correct reporter based on the flag
			stdoutReport := reporter.StdoutReport{Responses: responses}
			stdoutReport.Report()

			csvReport := reporter.CsvReport{Responses: responses}
			csvReport.Report()

			fmt.Println("Done!")
		},
	}
}
