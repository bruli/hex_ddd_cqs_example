package app_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/internal/app"
	"hex_ddd_cqs_example/internal/domain/user"
	"hex_ddd_cqs_example/internal/fixtures"
	"testing"
)

func TestCreateUser_Handle(t *testing.T) {
	errTest := errors.New("test error")
	us := fixtures.UserBuilder{}.Build(t)
	type args struct {
		cmd app.Command
	}
	validArgs := args{cmd: app.CreateUserCommand{
		ID:       us.Id(),
		UserName: us.UserName(),
		Phone:    us.Phone(),
	}}
	tests := []struct {
		name                 string
		args                 args
		expectedErr, findErr error
		user                 *user.User
	}{
		{
			name:        "with an invalid command, then it returns an invalid command error",
			args:        args{cmd: InvalidCommand{}},
			expectedErr: app.InvalidCommandError{},
		},
		{
			name:        "with a valid command when build user returns an error, then it returns a create user error",
			args:        args{cmd: app.CreateUserCommand{}},
			expectedErr: user.CreateUserError{},
		},
		{
			name:        "with a valid command when service return an error, then it returns same error",
			args:        validArgs,
			findErr:     errTest,
			expectedErr: errTest,
		},
	}
	for _, tt := range tests {
		t.Run(`Given a CreateUser command handler,
		when Handle method is called`+tt.name, func(t *testing.T) {
			t.Parallel()
			repo := &user.UserRepositoryMock{}
			repo.FindByIdFunc = func(ctx context.Context, id uuid.UUID) (*user.User, error) {
				return tt.user, tt.findErr
			}
			repo.SaveFunc = func(ctx context.Context, u user.User) error {
				return nil
			}
			handler := app.NewCreateUser(repo)
			_, err := handler.Handle(context.Background(), tt.args.cmd)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
		})
	}
}

type InvalidCommand struct{}

func (i InvalidCommand) Name() string {
	return "InvalidCommand"
}
