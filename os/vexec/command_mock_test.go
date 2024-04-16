package vexec

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockCommand_Run(t *testing.T) {
	// Create a new instance of MockCommand
	mockCommand := MockCommand{}

	// Test case 1: MockCommand execution with a valid command
	mockCommand.SetOutput([]byte("Mock output for command: echo Hello, World!"))
	output, err := mockCommand.Run("echo Hello, World!")
	assert.NoError(t, err)
	assert.Equal(t, "Mock output for command: echo Hello, World!", string(output))

	// Test case 2: MockCommand execution with an invalid command
	mockCommand.SetOutput([]byte("Mock output for command: invalidCommand"))
	output, err = mockCommand.Run("invalidCommand")
	assert.NoError(t, err)
	assert.Equal(t, "Mock output for command: invalidCommand", string(output))
}

func TestMockCommand_RunContext(t *testing.T) {
	// Create a new instance of MockCommand
	mockCommand := MockCommand{}

	// Test case 1: MockCommand execution with context and a valid command
	mockCommand.SetOutput([]byte("Mock output for command with context: echo Hello, World!"))
	output, err := mockCommand.RunContext(context.Background(), "echo Hello, World!")
	assert.NoError(t, err)
	assert.Equal(t, "Mock output for command with context: echo Hello, World!", string(output))

	// Test case 2: MockCommand execution with context and an invalid command
	mockCommand.SetOutput([]byte("Mock output for command with context: invalidCommand"))
	output, err = mockCommand.RunContext(context.Background(), "invalidCommand")
	assert.NoError(t, err)
	assert.Equal(t, "Mock output for command with context: invalidCommand", string(output))
}

func TestMockCommand_SetOutput(t *testing.T) {
	// Create a new instance of MockCommand
	mockCommand := MockCommand{}

	// Set output to "Test Output 1"
	mockCommand.SetOutput([]byte("Test Output 1"))

	// Test the Run method with the set output
	output, err := mockCommand.Run("testCommand1")
	assert.NoError(t, err)
	assert.Equal(t, "Test Output 1", string(output))

	// Set output to "Test Output 2"
	mockCommand.SetOutput([]byte("Test Output 2"))

	// Test the Run method with the new set output
	output, err = mockCommand.Run("testCommand2")
	assert.NoError(t, err)
	assert.Equal(t, "Test Output 2", string(output))
}
