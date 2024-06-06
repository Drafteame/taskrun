package models

import (
	"strings"

	"gopkg.in/yaml.v3"
)

var obfuscateEnvKeys = []string{
	"pass",
	"secret",
	"token",
	"aws",
	"user",
	"auth",
}

const obfuscatedValue = "********"

type JobConfig struct {
	Name       string   `yaml:"name"`
	WorkingDir string   `yaml:"working_dir,omitempty"`
	Command    string   `yaml:"command"`
	Args       []string `yaml:"args,omitempty"`
	Env        map[string]string
}

func (j JobConfig) ToYAML() string {
	by, err := yaml.Marshal(j.obfuscate())
	if err != nil {
		panic(err)
	}

	return string(by)
}

func (j JobConfig) obfuscate() JobConfig {
	if len(j.Args) == 0 {
		j.Args = nil
	}

	allArgs := strings.Join(j.Args, ",")

	for key, val := range j.Env {
		for _, obfuscateKey := range obfuscateEnvKeys {
			if strings.Contains(strings.ToLower(key), obfuscateKey) {
				j.Env[key] = obfuscatedValue
				j.Command = strings.ReplaceAll(j.Command, val, obfuscatedValue)
				allArgs = strings.ReplaceAll(allArgs, val, obfuscatedValue)
			}
		}
	}

	j.Args = strings.Split(allArgs, ",")

	return j
}
