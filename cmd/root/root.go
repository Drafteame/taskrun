package root

import (
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"

	"github.com/Drafteame/taskrun/internal"
	awsconfig "github.com/Drafteame/taskrun/internal/aws/config"
)

var rootCmd = &cobra.Command{
	Use:     "taskrun",
	Version: internal.Version,
	Short:   "Job execution tool tool",
	Long:    "Command line interface to automate task commands on configurable jobs",
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

	defaultJobsFile := wd + "/jobs-config.yml"

	rootCmd.PersistentFlags().StringVarP(&awsProfileFlag, "aws-profile", "", "", "AWS profile to use")
	rootCmd.PersistentFlags().StringVarP(&awsAccessKey, "aws-access-key", "", "", "AWS access key to use")
	rootCmd.PersistentFlags().StringVarP(&awsSecretKey, "aws-secret-key", "", "", "AWS secret key to use")
	rootCmd.PersistentFlags().StringVarP(&awsRegionFlag, "aws-region", "", "", "AWS region to use")
	rootCmd.PersistentFlags().StringVarP(&jobsFileFlag, "jobs-file", "j", defaultJobsFile, "Path to the jobs file")
	rootCmd.PersistentFlags().StringVarP(&stageFlag, "stage", "s", "", "Stage to run the migrations")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug mode")
}

func GetCommand() *cobra.Command {
	return rootCmd
}

func GetStage() *string {
	return &stageFlag
}

func GetJobsFile() *string {
	return &jobsFileFlag
}

func GetDebug() *bool {
	return &debugFlag
}

func GetReplacers() map[string]string {
	return map[string]string{
		`self(\s+)?:(\s+)?stage`: stageFlag,
		`sys(\s+)?:(\s+)?cwd`:    GetWorkingDir(),
	}
}

func GetAwsConfig() aws.Config {
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

func GetWorkingDir() string {
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
