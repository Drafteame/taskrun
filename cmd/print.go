package cmd

import (
	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/templating"
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
	RootCmd.AddCommand(printCmd)
}

func printJob(cmd *cobra.Command, args []string) {
	jobName := args[0]

	job, err := config.GetJob(jobName, stageFlag, jobsFileFlag)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	jobCfg, errRender := templating.NewJobTemplate(job).
		WithData(getReplacers()).
		WithAWSConfig(getAwsConfig()).
		Render()

	if errRender != nil {
		log.Fatal("Error: ", errRender)
	}

	printf("Config: \n---------\n%s", jobCfg.ToYAML())
}
