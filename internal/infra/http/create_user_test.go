package http_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/internal/app"
	"hex_ddd_cqs_example/internal/domain/user"
	"hex_ddd_cqs_example/internal/infra/http"
	http2 "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreteUser(t *testing.T) {
	validRequestBody := http.CreateUserRequest{
		ID:       uuid.New().String(),
		Username: "user test",
		Phone:    nil,
	}
	tests := []struct {
		name, body   string
		chErr        error
		expectedCode int
	}{
		{
			name:         "with an invalid body, then it returns a bad request",
			body:         "invalid",
			expectedCode: http2.StatusBadRequest,
		},
		{
			name:         "with a valid body but invalid id, then it returns a bad request",
			body:         buildRequestJsonToString(t, http.CreateUserRequest{}),
			expectedCode: http2.StatusBadRequest,
		},
		{
			name:         "with a valid body and command handler returns a create user error, then it returns a bad request",
			body:         buildRequestJsonToString(t, validRequestBody),
			chErr:        user.CreateUserError{},
			expectedCode: http2.StatusBadRequest,
		},
		{
			name:         "with a valid body and command handler returns an error, then it returns an internal server error",
			body:         buildRequestJsonToString(t, validRequestBody),
			chErr:        errors.New("error"),
			expectedCode: http2.StatusInternalServerError,
		},
		{
			name:         "with a valid body and command handler returns nil, then it returns ok",
			body:         buildRequestJsonToString(t, validRequestBody),
			expectedCode: http2.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(`Given a Create user endpoint
		when a request is sent`+tt.name, func(t *testing.T) {
			t.Parallel()
			ch := &app.CommandHandlerMock{}
			ch.HandleFunc = func(ctx context.Context, cmd app.Command) ([]app.Event, error) {
				return nil, tt.chErr
			}
			handler := http.CreteUser(ch)
			server := gin.Default()
			server.POST("/users", handler)

			req := httptest.NewRequest(http2.MethodPost, "/users", buildRequestBody(tt.body))
			writer := httptest.NewRecorder()
			server.ServeHTTP(writer, req)
			require.Equal(t, tt.expectedCode, writer.Code)
		})
	}
}

func buildRequestBody(body string) *strings.Reader {
	return strings.NewReader(body)
}

func buildRequestJsonToString(t *testing.T, req any) string {
	d, err := json.Marshal(req)
	require.NoError(t, err)
	return string(d)
}
