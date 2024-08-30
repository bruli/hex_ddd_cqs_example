package user

import (
	"context"
	"github.com/google/uuid"
)

func FindUser(ctx context.Context, repo UserRepository, id uuid.UUID) (*User, error) {
	return repo.FindById(ctx, id)
}
