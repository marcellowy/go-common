package vexec

import (
	"context"
	"os/exec"
)

// Command represents a command to be executed
type Command struct{}

// Run executes the command and returns the output and an error if any
func (c *Command) Run(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}

// RunContext executes the command with the given context and returns the output and an error if any
func (c *Command) RunContext(ctx context.Context, command string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
