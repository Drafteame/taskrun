package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"

	awsconfig "github.com/Drafteame/taskrun/internal/aws/config"
	"github.com/Drafteame/taskrun/internal/config"
	"github.com/Drafteame/taskrun/internal/models"
)

var RootCmd = &cobra.Command{
	Use:   "taskrun",
	Short: "Job execution tool tool",
	Long:  "Command line interface to automate task commands on configurable jobs",
}

var (
	jobsFileFlag   string
	stageFlag      string
	debugFlag      bool
	awsProfileFlag string
	awsRegionFlag  string
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defaultJobsFile := fmt.Sprintf("%s/jobs-config.yml", wd)

	RootCmd.PersistentFlags().StringVarP(&awsRegionFlag, "aws-region", "r", "", "AWS region to use")
	RootCmd.PersistentFlags().StringVarP(&awsProfileFlag, "aws-profile", "p", "", "AWS profile to use")
	RootCmd.PersistentFlags().StringVarP(&jobsFileFlag, "jobs-file", "j", defaultJobsFile, "Path to the jobs file")
	RootCmd.PersistentFlags().StringVarP(&stageFlag, "stage", "s", "", "Stage to run the migrations")
	RootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug mode")
}

func getReplacers() map[string]string {
	return map[string]string{
		"flag:stage": stageFlag,
		"sys:cwd":    getWorkingDir(),
	}
}

func setStage(cfg models.Jobs) {
	if stageFlag == "" {
		stageFlag = cfg.DefaultStage
	}

	if stageFlag == "" {
		for stage := range cfg.Jobs {
			stageFlag = stage
			break
		}
	}

	if stageFlag == "" {
		log.Fatal("Error: no stage defined")
	}
}

func getJobs() ([]models.Job, error) {
	cfg, err := config.LoadConfigFromPath(jobsFileFlag, getReplacers())
	if err != nil {
		return []models.Job{}, err
	}

	setStage(cfg)

	stageJobs, ok := cfg.Jobs[stageFlag]
	if !ok {
		return []models.Job{}, fmt.Errorf("stage %s not found", stageFlag)
	}

	return stageJobs, nil
}

func getJob(job string) (models.Job, error) {
	stageJobs, err := getJobs()
	if err != nil {
		return models.Job{}, err
	}

	for _, j := range stageJobs {
		if j.Name == job {
			return j, nil
		}
	}

	return models.Job{}, fmt.Errorf("job %s not found on stage %s", job, stageFlag)
}

func getAwsConfig() aws.Config {
	opts := make([]awsconfig.Option, 0, 2)

	if awsRegionFlag != "" {
		opts = append(opts, awsconfig.WithRegion(awsRegionFlag))
	}

	if awsProfileFlag != "" {
		opts = append(opts, awsconfig.WithProfile(awsProfileFlag))
	}

	cfg, err := awsconfig.Get(opts...)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	return cfg
}

func getWorkingDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	basePath := path.Dir(jobsFileFlag)

	if path.IsAbs(basePath) {
		cwd = basePath
	} else {
		cwd = path.Join(cwd, basePath)
	}

	return cwd
}

func printf(format string, a ...any) {
	_, _ = fmt.Printf(format, a...)
}
