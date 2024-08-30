package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
	"hex_ddd_cqs_example/internal/domain/user"
)

const UsersTableName = "users"

type UserRepository struct {
	sess db.Session
}

func NewUserRepository(sess db.Session) *UserRepository {
	return &UserRepository{sess: sess}
}

func (p UserRepository) Save(ctx context.Context, u user.User) error {
	if _, err := p.sess.WithContext(ctx).Collection(UsersTableName).Insert(u); err != nil {
		return err
	}
	return nil
}

func (p UserRepository) FindById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var model PostgresUser
	if err := p.sess.WithContext(ctx).Collection(UsersTableName).Find(db.Cond{"id": id.String()}).One(&model); err != nil {
		return nil, err
	}
	return buildUser(model)
}

func buildUser(model PostgresUser) (*user.User, error) {
	return user.NewUser(model.ID, model.UserName, model.Phone)
}

type PostgresUser struct {
	ID       uuid.UUID `db:"id"`
	UserName string    `db:"username"`
	Phone    *string   `db:"phone"`
}
