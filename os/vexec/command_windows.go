package vexec

import (
	"context"
	"os/exec"
)

// Command represents a command to be executed
type Command struct {
	Dir string
}

func NewCommand() *Command {
	return &Command{}
}

// Run executes the command and returns the output and an error if any
func (c *Command) Run(command string) ([]byte, error) {
	cmd := exec.Command("cmd", "/C", command)
	cmd.Dir = c.Dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}

// RunContext executes the command with the given context and returns the output and an error if any
func (c *Command) RunContext(ctx context.Context, command string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "cmd", "/C", command)
	cmd.Dir = c.Dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
