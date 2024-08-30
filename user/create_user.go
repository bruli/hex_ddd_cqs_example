package user

import (
	"context"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists")

// CreateUser NEW, injecting user repository and moving the database logic.
// now available to test.
func CreateUser(ctx context.Context, repo UserRepository, u User) error {
	current, err := FindUser(ctx, repo, u.ID)
	if err != nil {
		return err
	}
	if current != nil {
		return ErrUserAlreadyExists
	}

	return repo.Save(ctx, u)
}
