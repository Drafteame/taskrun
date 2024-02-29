package models

import (
	"gopkg.in/yaml.v3"
)

type JobConfig struct {
	Name       string   `yaml:"name"`
	WorkingDir string   `yaml:"working_dir,omitempty"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args,omitempty"`
	Env        map[string]string
}

func (j JobConfig) ToYAML() string {
	by, err := yaml.Marshal(j)
	if err != nil {
		panic(err)
	}

	return string(by)
}
