package models

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/Drafteame/taskrun/internal/aws/secretsmanager"
	"github.com/Drafteame/taskrun/internal/aws/ssm"
)

type Env struct {
	Remote EnvRemote         `yaml:"remote,omitempty"`
	Vars   map[string]EnvVar `yaml:"vars,omitempty"`
}

func (e Env) GetEnvs(cfg aws.Config) map[string]string {
	envs := make(map[string]string)

	for k, v := range e.Remote.GetValues(cfg) {
		envs[k] = v
	}

	for k, v := range e.Vars {
		envs[k] = v.GetValue(cfg)
	}

	return envs
}

type EnvRemote struct {
	Type string `yaml:"type"`
	Key  string `yaml:"key"`
}

func (e EnvRemote) GetValues(cfg aws.Config) map[string]string {
	if e.Type == "" {
		return map[string]string{}
	}

	switch e.Type {
	case "secretsmanager":
		return secretsmanager.MustGetSecret(e.Key, cfg)
	default:
		panic("Unknown remote source")
	}
}

type EnvVar struct {
	Source string `yaml:"source,omitempty"`
	Key    string `yaml:"key,omitempty"`
	Value  string `yaml:"value,omitempty"`
}

func (e EnvVar) GetValue(cfg aws.Config) string {
	switch e.Source {
	case "ssm":
		return ssm.MustGetParameter(e.Key, cfg)
	case "env":
		return os.Getenv(e.Key)
	default:
		return e.Value
	}
}
