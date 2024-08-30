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

func TestNewFindUser(t *testing.T) {
	us := fixtures.UserBuilder{}.Build(t)
	query := app.FindUserQuery{ID: us.Id()}
	errTest := errors.New("test error")
	type args struct {
		query app.Query
	}
	tests := []struct {
		name                 string
		args                 args
		expectedResult       any
		user                 *user.User
		repoErr, expectedErr error
	}{
		{
			name:        "with an invalid query, then it returns an invalid query error",
			args:        args{query: invalidQuery{}},
			expectedErr: app.InvalidQueryError{},
		},
		{
			name:        "with valid query and service returns an error, then it returns same error",
			args:        args{query: query},
			expectedErr: errTest,
			repoErr:     errTest,
		},
		{
			name:           "with valid query and service returns an user, then it returns valid result",
			args:           args{query: query},
			user:           &us,
			expectedResult: &us,
		},
	}
	for _, tt := range tests {
		t.Run(`Given a FindUser query handler,
		when Handle method is called`+tt.name, func(t *testing.T) {
			t.Parallel()
			repo := &user.UserRepositoryMock{}
			repo.FindByIdFunc = func(ctx context.Context, id uuid.UUID) (*user.User, error) {
				return tt.user, tt.repoErr
			}
			handler := app.NewFindUser(repo)
			result, err := handler.Handle(context.Background(), tt.args.query)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

type invalidQuery struct{}

func (i invalidQuery) Name() string {
	return "invalid query"
}
