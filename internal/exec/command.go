package exec

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

type Command struct {
	cmd    string
	args   []string
	env    map[string]string
	dir    string
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewCommand(cmd string, args ...string) *Command {
	return &Command{
		cmd:  cmd,
		args: args,
	}
}

func (c *Command) WithEnvs(env map[string]string) *Command {
	c.env = env
	return c
}

func (c *Command) WithWorkingDir(dir string) *Command {
	c.dir = dir
	return c
}

func (c *Command) Run() error {
	cmd := exec.Command(c.cmd, c.args...)

	cmd.Stdout = io.MultiWriter(os.Stdout, &c.stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &c.stderr)

	if c.dir != "" {
		cmd.Dir = c.dir
	}

	if c.env != nil {
		cmd.Env = os.Environ()
		for k, v := range c.env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}

	return cmd.Run()
}
