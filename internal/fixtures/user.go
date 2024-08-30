package fixtures

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/internal/domain/user"
	"testing"
)

type UserBuilder struct {
	ID       *uuid.UUID
	Username *string
	Phone    *string
}

func (b UserBuilder) Build(t *testing.T) user.User {
	id := setData(uuid.New(), b.ID)
	username := setData("username", b.Username)
	u, err := user.NewUser(id, username, b.Phone)
	require.NoError(t, err)
	return *u
}

func setData[T any](def T, value *T) T {
	if value != nil {
		return *value
	}
	return def
}
