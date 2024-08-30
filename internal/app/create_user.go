package app

import (
	"context"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/internal/domain/user"
)

const CreateUserCommandName = "createUser"

type CreateUserCommand struct {
	ID       uuid.UUID
	UserName string
	Phone    *string
}

func (c CreateUserCommand) Name() string {
	return CreateUserCommandName
}

type CreateUser struct {
	repo user.UserRepository
}

func (c CreateUser) Handle(ctx context.Context, cmd Command) ([]Event, error) {
	co, ok := cmd.(CreateUserCommand)
	if !ok {
		return nil, NewInvalidCommandError(cmd.Name(), CreateUserCommandName)
	}
	us, err := user.NewUser(co.ID, co.UserName, co.Phone)
	if err != nil {
		return nil, user.NewCreateUserError(err.Error())
	}
	return nil, user.CreateUser(ctx, c.repo, *us)
}

func NewCreateUser(repo user.UserRepository) *CreateUser {
	return &CreateUser{repo: repo}
}
