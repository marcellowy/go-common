package vexec

func TestCommand_Run(t *testing.T) {
	// Create a new instance of Command
	command := Command{}

	// Test case 1: Command execution with a valid command
	output, err := command.Run("echo Hello, World!")
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!\r\n", string(output))

	// Test case 2: Command execution with an invalid command
	output, err = command.Run("invalidCommand")
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestCommand_RunContext(t *testing.T) {
	// Create a new instance of Command
	command := Command{}

	// Test case 1: Command execution with context and a valid command
	output, err := command.RunContext(context.Background(), "echo Hello, World!")
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!\r\n", string(output))

	// Test case 2: Command execution with context and an invalid command
	output, err = command.RunContext(context.Background(), "invalidCommand")
	assert.Error(t, err)
	assert.Nil(t, output)
}
