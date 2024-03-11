package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/config"
)

var listTasksCmd = &cobra.Command{
	Use:   "list-tasks",
	Short: "List all available tasks",
	Long:  "List all available tasks for database migration",
	Run:   listTasks,
}

func init() {
	RootCmd.AddCommand(listTasksCmd)
}

func listTasks(cmd *cobra.Command, args []string) {
	stageJobs, err := config.GetJobs(stageFlag, jobsFileFlag)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	printf("Jobs for stage %s:\n---------\n", stageFlag)

	for _, j := range stageJobs {
		printf("- %s\n", j.Name)
	}
}
