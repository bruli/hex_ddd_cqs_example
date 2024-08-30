package app

import "context"

type Command interface {
	Name() string
}

//go:generate moq -out zmock_command_handler.go . CommandHandler
type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) ([]Event, error)
}
