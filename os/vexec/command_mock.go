package vexec

import "context"

// MockCommand is a mock implementation of Command for testing
type MockCommand struct {
	output []byte
}

// NewMockCommand creates a new instance of MockCommand with a default output
func NewMockCommand() *MockCommand {
	return &MockCommand{
		output: []byte("Default mock output"),
	}
}

// SetOutput sets the desired output for the mock command execution
func (c *MockCommand) SetOutput(output []byte) {
	c.output = output
}

// Run simulates executing a command and returns the set output
func (c *MockCommand) Run(command string) ([]byte, error) {
	// Simulate the execution of the command and return the set output
	return c.output, nil
}

// RunContext simulates executing a command with context and returns the set output
func (c *MockCommand) RunContext(ctx context.Context, command string) ([]byte, error) {
	// Simulate the execution of the command with context and return the set output
	return c.output, nil
}
