package vexec

import "context"

// InterfaceCommand is an interface for running commands
type InterfaceCommand interface {
	Run(command string) ([]byte, error)
	RunContext(ctx context.Context, command string) ([]byte, error)
}
