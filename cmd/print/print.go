package print

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/cmd/root"
	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/console"
	"github.com/Drafteame/taskrun/internal/templating"
)

var printCmd = &cobra.Command{
	Use:   "print [job-name]",
	Short: "Print a job",
	Long:  "Print the requested job configuration",
	Args:  cobra.ExactArgs(1),
	Run:   printJob,
}

var (
	stageFlag    *string
	jobsFileFlag *string
)

func GetCommand(stage, jobsFile *string) *cobra.Command {
	stageFlag = stage
	jobsFileFlag = jobsFile

	return printCmd
}

func printJob(cmd *cobra.Command, args []string) {
	jobName := args[0]

	job, err := config.GetJob(jobName, *stageFlag, *jobsFileFlag)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	jobCfg, errRender := templating.NewJobTemplate(job).
		WithData(root.GetReplacers()).
		WithAWSConfig(root.GetAwsConfig()).
		Render()

	if errRender != nil {
		log.Fatal("Error: ", errRender)
	}

	console.Printf("Config: \n---------\n%s", jobCfg.ToYAML())
}
