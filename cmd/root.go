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
	RootCmd.PersistentFlags().StringVarP(&stageFlag, "stage", "s", "local", "Stage to run the migrations")
	RootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug mode")
}

func getReplacers() map[string]string {
	return map[string]string{
		"flag:stage": stageFlag,
		"sys:cwd":    getWorkingDir(),
	}
}

func getJob(job string) (models.Job, error) {
	cfg, err := config.LoadConfigFromPath(jobsFileFlag, getReplacers())
	if err != nil {
		return models.Job{}, err
	}

	stageJobs, ok := cfg.Jobs[stageFlag]
	if !ok {
		return models.Job{}, fmt.Errorf("stage %s not found", stageFlag)
	}

	for _, j := range stageJobs {
		if j.Name == job {
			return j, nil
		}
	}

	return models.Job{}, fmt.Errorf("job %s not found", job)
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

	basepath := path.Dir(jobsFileFlag)

	if path.IsAbs(basepath) {
		cwd = basepath
	} else {
		cwd = path.Join(cwd, basepath)
	}

	return cwd
}

func setWorkingDir() {
	if err := os.Chdir(getWorkingDir()); err != nil {
		panic(err)
	}
}
