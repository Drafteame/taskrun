package models

type Env struct {
	Remote EnvRemote         `yaml:"remote,omitempty"`
	Vars   map[string]EnvVar `yaml:"vars,omitempty"`
}

func (e Env) HasVars() bool {
	return len(e.Vars) > 0
}

func (e Env) HasRemote() bool {
	return e.Remote.Type != ""
}

type EnvRemote struct {
	Type string `yaml:"type"`
	Key  string `yaml:"key"`
}

type EnvVar struct {
	Source string `yaml:"source,omitempty"`
	Key    string `yaml:"key,omitempty"`
	Value  string `yaml:"value,omitempty"`
}
