package run

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/cmd/root"
	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/console"
	"github.com/Drafteame/taskrun/internal/exec"
	"github.com/Drafteame/taskrun/internal/models"
	"github.com/Drafteame/taskrun/internal/templating"
)

var runJobCmd = &cobra.Command{
	Use:   "run [task-name]",
	Short: "Run a task",
	Long:  "Run the requested task configuration",
	Args:  cobra.ExactArgs(1),
	Run:   runJob,
}

var (
	stageFlag    *string
	jobsFileFlag *string
	debugFlag    *bool
)

func GetCommand(stage, jobsFile *string, debug *bool) *cobra.Command {
	stageFlag = stage
	jobsFileFlag = jobsFile
	debugFlag = debug

	return runJobCmd
}

func runJob(cmd *cobra.Command, args []string) {
	jobName := args[0]

	job, err := config.GetJob(jobName, *stageFlag, *jobsFileFlag)
	if err != nil {
		log.Fatal("Error getting job: ", err)
	}

	jobCfg, errRender := templating.NewJobTemplate(job).
		WithData(root.GetReplacers()).
		WithAWSConfig(root.GetAwsConfig()).
		Render()

	if errRender != nil {
		log.Fatal("Error: ", errRender)
	}

	if debugFlag != nil && *debugFlag {
		console.Printf("Config: \n---------\n%s", jobCfg.ToYAML())
	}

	executeCommand(jobCfg)
}

func executeCommand(j *models.JobConfig) {
	cwd := root.GetWorkingDir()

	println("Running job: ", j.Name)
	println("Working dir: ", cwd)

	cmd := exec.NewCommand(j.Command, j.Args...).
		WithEnvs(j.Env).
		WithWorkingDir(cwd)

	if err := cmd.Run(); err != nil {
		log.Fatal("Error: ", err)
	}
}
