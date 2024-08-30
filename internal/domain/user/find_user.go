package user

import (
	"context"
	"github.com/google/uuid"
)

// FindUser NEW, injecting repository and moving the database login into database object
// available to test
func FindUser(ctx context.Context, repo UserRepository, id uuid.UUID) (*User, error) {
	return repo.FindById(ctx, id)
}
