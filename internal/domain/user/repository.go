package user

import (
	"context"
	"github.com/google/uuid"
)

//go:generate moq -out zmock_repository_test.go . UserRepository
type UserRepository interface {
	Save(ctx context.Context, u User) error
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
}
