package user

import (
	"context"
	"errors"
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
