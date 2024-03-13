package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/models"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	Long:  "List all available tasks for database migration",
	Run:   listTasks,
}

var ltJSONFlag bool

func init() {
	listCmd.Flags().BoolVarP(&ltJSONFlag, "json", "", false, "Output as JSON")

	RootCmd.AddCommand(listCmd)
}

func listTasks(cmd *cobra.Command, args []string) {
	jobs, err := config.GetJobs(stageFlag, jobsFileFlag)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	if ltJSONFlag {
		printTasksAsJSON(jobs)
		return
	}

	printTasks(jobs)
}

func printTasks(jobs []models.Job) {
	printf("Jobs for stage %s:\n---------\n", stageFlag)

	for _, j := range jobs {
		printf("- %s\n", j.Name)
	}
}

func printTasksAsJSON(jobs []models.Job) {
	names := make([]string, 0, len(jobs))

	for _, j := range jobs {
		names = append(names, j.Name)
	}

	jb, err := json.Marshal(names)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	println(string(jb))
}
