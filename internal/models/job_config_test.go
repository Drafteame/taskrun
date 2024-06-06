package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestJobConfigToYAML(t *testing.T) {
	type tt struct {
		name string
		jc   JobConfig
		want string
	}

	tests := []tt{
		{
			name: "empty",
			jc:   JobConfig{},
			want: `name: ""
command: ""
args:
    - ""
env: {}
`,
		},
		{
			name: "full",
			jc: JobConfig{
				Name:    "test",
				Command: "echo 'hello world'",
				Args: []string{
					"hello",
					"world",
				},
				Env: map[string]string{
					"key": "value",
				},
			},
			want: `name: test
command: echo 'hello world'
args:
    - hello
    - world
env:
    key: value
`,
		},
		{
			name: "obfuscated password",
			jc: JobConfig{
				Name:    "test",
				Command: "echo 'mypassword is: password'",
				Env: map[string]string{
					"password": "password",
				},
			},
			want: `name: test
command: 'echo ''my******** is: ********'''
args:
    - ""
env:
    password: '********'
`,
		},
		{
			name: "obfuscated secret",
			jc: JobConfig{
				Name:    "test",
				Command: "echo 'my value is: secret'",
				Env: map[string]string{
					"secret": "secret",
				},
			},
			want: `name: test
command: 'echo ''my value is: ********'''
args:
    - ""
env:
    secret: '********'
`,
		},
		{
			name: "obfuscated aws envs",
			jc: JobConfig{
				Name:    "test",
				Command: "echo 'my value is: aws_value aws_secret aws_session aws_region'",
				Env: map[string]string{
					"AWS_ACCESS_KEY_ID":     "aws_value",
					"aws_secret_access_key": "aws_secret",
					"aws_session_token":     "aws_session",
					"aws_region":            "aws_region",
				},
			},
			want: `name: test
command: 'echo ''my value is: ******** ******** ******** ********'''
args:
    - ""
env:
    AWS_ACCESS_KEY_ID: '********'
    aws_region: '********'
    aws_secret_access_key: '********'
    aws_session_token: '********'
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.jc.ToYAML()

			assert.Equal(t, tc.want, got)
		})
	}
}
