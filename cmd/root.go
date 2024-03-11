package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"

	awsconfig "github.com/Drafteame/taskrun/internal/aws/config"
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
	awsAccessKey   string
	awsSecretKey   string
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defaultJobsFile := fmt.Sprintf("%s/jobs-config.yml", wd)

	RootCmd.PersistentFlags().StringVarP(&awsProfileFlag, "aws-profile", "", "", "AWS profile to use")
	RootCmd.PersistentFlags().StringVarP(&awsAccessKey, "aws-access-key", "", "", "AWS access key to use")
	RootCmd.PersistentFlags().StringVarP(&awsSecretKey, "aws-secret-key", "", "", "AWS secret key to use")
	RootCmd.PersistentFlags().StringVarP(&awsRegionFlag, "aws-region", "", "", "AWS region to use")
	RootCmd.PersistentFlags().StringVarP(&jobsFileFlag, "jobs-file", "j", defaultJobsFile, "Path to the jobs file")
	RootCmd.PersistentFlags().StringVarP(&stageFlag, "stage", "s", "", "Stage to run the migrations")
	RootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug mode")
}

func getReplacers() map[string]string {
	return map[string]string{
		`self(\s+)?:(\s+)?stage`: stageFlag,
		`sys(\s+)?:(\s+)?cwd`:    getWorkingDir(),
	}
}

func getAwsConfig() aws.Config {
	opts := make([]awsconfig.Option, 0, 2)

	if awsRegionFlag != "" {
		opts = append(opts, awsconfig.WithRegion(awsRegionFlag))
	}

	if awsProfileFlag != "" {
		opts = append(opts, awsconfig.WithProfile(awsProfileFlag))
	}

	if awsAccessKey != "" {
		opts = append(opts, awsconfig.WithAccessKey(awsAccessKey))
	}

	if awsSecretKey != "" {
		opts = append(opts, awsconfig.WithSecretKey(awsSecretKey))
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
