package app

import "context"

type Query interface {
	Name() string
}

//go:generate moq -out zmock_query_handler.go . QueryHandler
type QueryHandler interface {
	Handle(ctx context.Context, query Query) (any, error)
}
