package cmd

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/mfbmina/enxame/reporter"
	"github.com/mfbmina/enxame/swarm"
	"github.com/spf13/cobra"
)

var format string
var output string
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

	rootCmd.AddCommand(runCmd())
}

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a test",
		Long:  "Run a test against the given URI",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Swarming", args[0], "...")
			responses := swarm.Swarm(args[0], requestsPerUser, users)

			fmt.Printf("Reporting results as %s...\n", format)
			r := reporter.NewReporter(format, output, responses)
			feedback, err := r.Report()
			if err != nil {
				sentry.CaptureException(err)
				return
			}

			fmt.Println(feedback)
			fmt.Println("Done!")
		},
	}

	cmd.Flags().IntVarP(&requestsPerUser, "requests_per_user", "r", 5, "Max number of requests per user")
	cmd.Flags().IntVarP(&users, "users", "u", 10, "Max number of concurrent users")
	cmd.Flags().StringVarP(&format, "format", "f", "txt", "The report format (txt, csv, json)")
	cmd.Flags().StringVarP(&output, "output", "o", "", "The output file name (extension will be appended automatically)")

	return cmd
}
