package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "List all available jobs",
	Long:  "List all available jobs for database migration",
	Run:   listJobs,
}

func init() {
	RootCmd.AddCommand(listJobsCmd)
}

func listJobs(cmd *cobra.Command, args []string) {
	stageJobs, err := getJobs()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	printf("Jobs for stage %s:\n---------\n", stageFlag)

	for _, j := range stageJobs {
		_, _ = fmt.Printf("- %s\n", j.Name)
	}
}
