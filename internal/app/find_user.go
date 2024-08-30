package app

import (
	"context"
	"github.com/google/uuid"
	"hex_ddd_cqs_example/internal/domain/user"
)

const FindUserQueryName = "FindUser"

type FindUserQuery struct {
	ID uuid.UUID
}

func (f FindUserQuery) Name() string {
	return FindUserQueryName
}

type FindUser struct {
	repo user.UserRepository
}

func (f FindUser) Handle(ctx context.Context, query Query) (any, error) {
	q, ok := query.(FindUserQuery)
	if !ok {
		return nil, NewInvalidQueryError(query.Name(), FindUserQueryName)
	}
	return user.FindUser(ctx, f.repo, q.ID)
}

func NewFindUser(repo user.UserRepository) *FindUser {
	return &FindUser{repo: repo}
}
