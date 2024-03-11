package cmd

import (
	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/templating"
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal/exec"
	"github.com/Drafteame/taskrun/internal/models"
)

var runJobCmd = &cobra.Command{
	Use:   "run [task-name]",
	Short: "Run a task",
	Long:  "Run the requested task configuration",
	Args:  cobra.ExactArgs(1),
	Run:   runJob,
}

func init() {
	RootCmd.AddCommand(runJobCmd)
}

func runJob(cmd *cobra.Command, args []string) {
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

	if debugFlag {
		printf("Config: \n---------\n%s", jobCfg.ToYAML())
	}

	executeCommand(jobCfg)
}

func executeCommand(j *models.JobConfig) {
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
