package models

import (
	"gopkg.in/yaml.v3"
)

type Job struct {
	Name       string   `yaml:"name"`
	WorkingDir string   `yaml:"working_dir,omitempty"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args,omitempty"`
	Env        Env      `yaml:"env,omitempty"`
}

func (j *Job) ToYAML() string {
	by, err := yaml.Marshal(j)
	if err != nil {
		panic(err)
	}

	return string(by)
}

func (j *Job) FromYAML(by []byte) error {
	return yaml.Unmarshal(by, &j)
}

func (j *Job) ToJobConfig(envs map[string]string) *JobConfig {
	return &JobConfig{
		Name:       j.Name,
		WorkingDir: j.WorkingDir,
		Command:    j.Command,
		Args:       j.Args,
		Env:        envs,
	}
}
