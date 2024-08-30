package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

var ErrUserAlreadyExists = errors.New("user already exists")

const UsersTableName = "users"

func CreateUser(ctx context.Context, session db.Session, u PostgresUser) error {
	current, err := FindUser(ctx, session, u.ID)
	if err != nil {
		return err
	}
	if current != nil {
		return ErrUserAlreadyExists
	}
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
