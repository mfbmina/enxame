package cmd

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/mfbmina/enxame/reporter"
	"github.com/mfbmina/enxame/swarm"
	"github.com/spf13/cobra"
)

var format string
var method string
var output string
var requestsPerUser int
var users int

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a test",
		Long:  "Run a test against the given URI",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Validating method...")
			fmt.Println("Method:", method)
			if !validMethod(method) {
				err := fmt.Errorf("Invalid method. Please use one of the following: GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD")
				sentry.CaptureException(err)
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Swarming", args[0], "...")
			responses, err := swarm.Swarm(args[0], method, requestsPerUser, users)
			if err != nil {
				sentry.CaptureException(err)
				return err
			}

			fmt.Printf("Reporting results as %s...\n", format)
			r := reporter.NewReporter(format, output, responses)
			feedback, err := r.Report()
			if err != nil {
				sentry.CaptureException(err)
				return err
			}

			fmt.Println(feedback)
			fmt.Println("Done!")
			return nil
		},
	}

	cmd.Flags().IntVarP(&requestsPerUser, "requests_per_user", "r", 5, "Max number of requests per user")
	cmd.Flags().IntVarP(&users, "users", "u", 10, "Max number of concurrent users")
	cmd.Flags().StringVarP(&format, "format", "f", "txt", "The report format (txt, csv, json)")
	cmd.Flags().StringVarP(&method, "method", "X", "GET", "The http method to use (GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD)")
	cmd.Flags().StringVarP(&output, "output", "o", "", "The output file name (extension will be appended automatically)")

	return cmd
}

func validMethod(method string) bool {
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
	f := false

	for _, m := range allowedMethods {
		if m == method {
			f = true
			break
		}
	}

	return f
}
