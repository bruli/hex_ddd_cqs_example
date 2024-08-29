package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const UsersTableName = "users"

func CreateUser(ctx context.Context, session db.Session, u PostgresUser) error {
	if _, err := session.WithContext(ctx).Collection(UsersTableName).Insert(u); err != nil {
		return err
	}
	return nil
}

type PostgresUser struct {
	ID       uuid.UUID `db:"id"`
	UserName string    `db:"username"`
	Phone    *string   `db:"phone"`
}
