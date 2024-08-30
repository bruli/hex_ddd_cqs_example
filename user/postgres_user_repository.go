package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const UsersTableName = "users"

type PostgresUserRepository struct {
	sess db.Session
}

func NewPostgresUserRepository(sess db.Session) *PostgresUserRepository {
	return &PostgresUserRepository{sess: sess}
}

func (p PostgresUserRepository) Save(ctx context.Context, u User) error {
	if _, err := p.sess.WithContext(ctx).Collection(UsersTableName).Insert(u); err != nil {
		return err
	}
	return nil
}

func (p PostgresUserRepository) FindById(ctx context.Context, id uuid.UUID) (*User, error) {
	var model PostgresUser
	if err := p.sess.WithContext(ctx).Collection(UsersTableName).Find(db.Cond{"id": id.String()}).One(&model); err != nil {
		return nil, err
	}
	return buildUser(model), nil
}

func buildUser(model PostgresUser) *User {
	return &User{
		ID:       model.ID,
		UserName: model.UserName,
		Phone:    model.Phone,
	}
}

type PostgresUser struct {
	ID       uuid.UUID `db:"id"`
	UserName string    `db:"username"`
	Phone    *string   `db:"phone"`
}
