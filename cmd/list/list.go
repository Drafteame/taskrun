package list

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/console"
	"github.com/Drafteame/taskrun/internal/models"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	Long:  "List all available tasks for database migration",
	Run:   listTasks,
}

var (
	ltJSONFlag   bool
	stageFlag    *string
	jobsFileFlag *string
)

func init() {
	listCmd.Flags().BoolVarP(&ltJSONFlag, "json", "", false, "Output as JSON")
}

func GetCommand(stage, jobsFile *string) *cobra.Command {
	stageFlag = stage
	jobsFileFlag = jobsFile

	return listCmd
}

func listTasks(cmd *cobra.Command, args []string) {
	jobs, err := config.GetJobs(*stageFlag, *jobsFileFlag)
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
	console.Printf("Jobs for stage %s:\n---------\n", *stageFlag)

	for _, j := range jobs {
		console.Printf("- %s\n", j.Name)
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

	_, _ = fmt.Fprint(os.Stdout, string(jb))
}
