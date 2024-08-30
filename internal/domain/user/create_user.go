package user

import (
	"context"
	"errors"
	"fmt"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func CreateUser(ctx context.Context, repo UserRepository, u User) error {
	current, err := FindUser(ctx, repo, u.Id())
	if err != nil {
		return err
	}
	if current != nil {
		return ErrUserAlreadyExists
	}

	return repo.Save(ctx, u)
}

type CreateUserError struct {
	m string
}

func NewCreateUserError(msg string) *CreateUserError {
	return &CreateUserError{m: msg}
}

func (c CreateUserError) Error() string {
	return fmt.Sprintf("failed to create user: %s", c.m)
}
