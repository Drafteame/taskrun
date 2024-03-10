package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/exec"
	"github.com/Drafteame/taskrun/internal/models"
)

var runJobCmd = &cobra.Command{
	Use:   "run [job-name]",
	Short: "Run a job",
	Long:  "Run the requested job configuration",
	Args:  cobra.ExactArgs(1),
	Run:   runJob,
}

func init() {
	RootCmd.AddCommand(runJobCmd)
}

func runJob(cmd *cobra.Command, args []string) {
	jobName := args[0]

	job, err := getJob(jobName)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	jobConfig := job.ToJobConfig(getAwsConfig())

	if debugFlag {
		printf("Config: \n---------\n%s", jobConfig.ToYAML())
	}

	executeCommand(jobConfig)
}

func executeCommand(j models.JobConfig) {
	cwd := getWorkingDir()

	println("Running job: ", j.Name)
	println("Working dir: ", cwd)

	cmd := exec.NewCommand(j.Command, j.Args...).
		WithEnvs(j.Env).
		WithWorkingDir(cwd)

	if err := cmd.Run(); err != nil {
		log.Fatal("Error: ", err)
	}
}
