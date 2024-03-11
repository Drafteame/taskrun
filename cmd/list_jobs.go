package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/config"
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
	stageJobs, err := config.GetJobs(stageFlag, jobsFileFlag)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	printf("Jobs for stage %s:\n---------\n", stageFlag)

	for _, j := range stageJobs {
		printf("- %s\n", j.Name)
	}
}
