package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

func FindUser(ctx context.Context, session db.Session, id uuid.UUID) (*PostgresUser, error) {
	var model PostgresUser
	if err := session.WithContext(ctx).Collection(UsersTableName).Find(db.Cond{"id": id.String()}).One(&model); err != nil {
		return nil, err
	}
	return &model, nil
}
