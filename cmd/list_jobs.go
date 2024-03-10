package cmd

import (
	"fmt"
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
	cfg, err := config.LoadConfigFromPath(jobsFileFlag, getReplacers())
	if err != nil {
		log.Fatal("Error: ", err)
	}

	stageJobs, ok := cfg.Jobs[stageFlag]
	if !ok {
		log.Fatalf("Error: stage %s not found", stageFlag)
	}

	_, _ = fmt.Printf("Jobs for stage %s:\n---------\n", stageFlag)

	for _, j := range stageJobs {
		_, _ = fmt.Printf("- %s\n", j.Name)
	}
}
