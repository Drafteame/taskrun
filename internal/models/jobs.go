package models

import (
	"gopkg.in/yaml.v3"
)

type Jobs struct {
	DefaultStage string           `yaml:"default_stage"`
	Jobs         map[string][]Job `yaml:"jobs"`
}

func (j *Jobs) ToYAML() string {
	by, err := yaml.Marshal(j)
	if err != nil {
		panic(err)
	}

	return string(by)
}

func (j *Jobs) FromYAML(by []byte) error {
	return yaml.Unmarshal(by, j)
}
