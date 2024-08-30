package http_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"hex_ddd_cqs_example/internal/app"
	"hex_ddd_cqs_example/internal/fixtures"
	"hex_ddd_cqs_example/internal/infra/http"
	http2 "net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUser(t *testing.T) {
	us := fixtures.UserBuilder{}.Build(t)
	tests := []struct {
		name         string
		id           string
		qhResult     any
		qhErr        error
		expectedCode int
	}{
		{
			name:         "with an invalid user id in path, then it returns bad request",
			id:           "1",
			expectedCode: http2.StatusBadRequest,
		},
		{
			name:         "with a valid user id in path and query returns an error, then it returns internal server error",
			id:           us.Id().String(),
			qhErr:        errors.New("test error"),
			expectedCode: http2.StatusInternalServerError,
		},
		{
			name:         "with a valid user id in path and query returns an invalid result, then it returns internal server error",
			id:           us.Id().String(),
			qhResult:     "invalid",
			expectedCode: http2.StatusInternalServerError,
		},
		{
			name:         "with a valid user id in path and query returns a valid result, then it returns valid response",
			id:           us.Id().String(),
			qhResult:     &us,
			expectedCode: http2.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(`Given FindUser endpoint,
		when a request is sent`+tt.name, func(t *testing.T) {
			t.Parallel()
			qh := &app.QueryHandlerMock{}
			qh.HandleFunc = func(ctx context.Context, query app.Query) (any, error) {
				return tt.qhResult, tt.qhErr
			}
			handler := http.FindUser(qh)
			server := gin.Default()
			server.GET("/users/:id", handler)
			req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", tt.id), nil)
			writer := httptest.NewRecorder()
			server.ServeHTTP(writer, req)
			require.Equal(t, tt.expectedCode, writer.Code)
		})
	}
}
