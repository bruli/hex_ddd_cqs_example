package user_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/internal/domain/user"
	"hex_ddd_cqs_example/internal/fixtures"
	"testing"
)

func TestNewUser(t *testing.T) {
	us := fixtures.UserBuilder{}.Build(t)
	type args struct {
		id       uuid.UUID
		userName string
		phone    *string
	}
	tests := []struct {
		name        string
		args        args
		want        *user.User
		expectedErr error
	}{
		{
			name:        "with an invalid id, then it returns an invalid user id error",
			args:        args{},
			expectedErr: user.ErrInvalidUserID,
		},
		{
			name: "with an invalid username, then it returns an invalid user name error",
			args: args{
				id: us.Id(),
			},
			expectedErr: user.ErrInvalidUserName,
		},
		{
			name: "with all valid data, then it returns a valid user",
			args: args{
				id:       us.Id(),
				userName: us.UserName(),
				phone:    us.Phone(),
			},
			want: &us,
		},
	}
	for _, tt := range tests {
		t.Run(`Given a NewUser method,
		when is called `+tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewUser(tt.args.id, tt.args.userName, tt.args.phone)
			if err != nil {
				require.ErrorIs(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.Id(), got.Id())
			require.Equal(t, tt.want.UserName(), got.UserName())
			require.Equal(t, tt.want.Phone(), got.Phone())
		})
	}
}
