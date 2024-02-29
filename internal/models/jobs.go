package models

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"gopkg.in/yaml.v3"
)

type Jobs struct {
	Jobs map[string][]Job `yaml:"jobs"`
}

type Job struct {
	Name       string   `yaml:"name"`
	WorkingDir string   `yaml:"working_dir,omitempty"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args,omitempty"`
	Env        Env      `yaml:"env,omitempty"`
}

func (j Job) ToYAML() string {
	by, err := yaml.Marshal(j)
	if err != nil {
		panic(err)
	}

	return string(by)
}

func (j Job) ToJobConfig(cfg aws.Config) JobConfig {
	jc := JobConfig{
		Name:       j.Name,
		WorkingDir: j.WorkingDir,
		Command:    j.Command,
		Args:       j.Args,
	}

	envs := j.Env.GetEnvs(cfg)

	strConf := jc.ToYAML()

	for env, val := range envs {
		exp, err := regexp.Compile(fmt.Sprintf(`\$\{(\s+)?env:%s(\s+)?\}`, env))
		if err != nil {
			panic(err)
		}

		strConf = exp.ReplaceAllString(strConf, val)
	}

	exp, err := regexp.Compile(`\$\{(\s+)?env:(\w+)(\s+)?\}`)
	if err != nil {
		panic(err)
	}

	strConf = exp.ReplaceAllString(strConf, "")

	if err := yaml.Unmarshal([]byte(strConf), &jc); err != nil {
		panic(err)
	}

	jc.Env = envs

	return jc
}
