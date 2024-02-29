package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print [job-name]",
	Short: "Print a job",
	Long:  "Print the requested job configuration",
	Args:  cobra.ExactArgs(1),
	Run:   printJob,
}

func init() {
	rootCmd.AddCommand(printCmd)
}

func printJob(cmd *cobra.Command, args []string) {
	jobName := args[0]

	job, err := getJob(jobName)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	_, _ = fmt.Printf("Config: \n---------\n%s", job.ToJobConfig(getAwsConfig()).ToYAML())
}
