package user_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/user"
	"testing"
)

func TestCreateUser(t *testing.T) {
	errTest := errors.New("test error")
	type args struct {
		u user.User
	}
	u := user.User{}
	defaultArgs := args{u: u}
	tests := []struct {
		name                              string
		args                              args
		findRepoErr, saveErr, expectedErr error
		foundUser                         *user.User
	}{
		{
			name:        "and find user returns an error, then it returns same error",
			args:        defaultArgs,
			findRepoErr: errTest,
			expectedErr: errTest,
		},
		{
			name:        "and find user returns an existent user, then it returns user already exists error",
			args:        defaultArgs,
			foundUser:   &u,
			expectedErr: user.ErrUserAlreadyExists,
		},
		{
			name:        "and save user returns an error, then it returns same error",
			args:        defaultArgs,
			foundUser:   nil,
			saveErr:     errTest,
			expectedErr: errTest,
		},
		{
			name:      "and save user returns nil, then it returns nil",
			args:      defaultArgs,
			foundUser: nil,
		},
	}
	for _, tt := range tests {
		t.Run(`Given a CreateUser method,
		when is called `+tt.name, func(t *testing.T) {
			repo := &user.UserRepositoryMock{}
			repo.FindByIdFunc = func(ctx context.Context, id uuid.UUID) (*user.User, error) {
				return tt.foundUser, tt.findRepoErr
			}
			repo.SaveFunc = func(ctx context.Context, u user.User) error {
				return tt.saveErr
			}
			err := user.CreateUser(context.Background(), repo, tt.args.u)
			if err != nil {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
