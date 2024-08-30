package user_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/user"
	"testing"
)

func TestFindUser(t *testing.T) {
	errTest := errors.New("test error")
	type args struct {
		id uuid.UUID
	}
	id := uuid.New()
	defaultArgs := args{id: id}
	tests := []struct {
		name                 string
		args                 args
		want                 *user.User
		repoErr, expectedErr error
	}{
		{
			name:        "and repository returns an error, then it returns same error",
			args:        defaultArgs,
			repoErr:     errTest,
			expectedErr: errTest,
		},
		{
			name: "and repository returns an user, then it returns same user",
			args: defaultArgs,
			want: &user.User{ID: id},
		},
	}
	for _, tt := range tests {
		t.Run(`Given a FindUser method,
		when is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			repo := &user.UserRepositoryMock{}
			repo.FindByIdFunc = func(ctx context.Context, id uuid.UUID) (*user.User, error) {
				return tt.want, tt.repoErr
			}
			got, err := user.FindUser(context.Background(), repo, tt.args.id)
			if err != nil {
				require.Equal(t, tt.expectedErr, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
