package cmd

import (
	"fmt"

	"github.com/mfbmina/enxame/lib"
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

			for _, r := range lib.Swarm(args[0], requestsPerUser, users) {
				fmt.Println(r.StatusCode, r.Time, r.Path)
			}

			fmt.Println("Done!")
		},
	}
}
