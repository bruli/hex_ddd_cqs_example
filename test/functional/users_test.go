//go:build functional

package functional

import (
	"fmt"
	"github.com/google/uuid"
	http2 "hex_ddd_cqs_example/http"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	userID := uuid.New()
	t.Run(`Given a create user endpoint,
	when a request is sent,
	then it returns an OK status`, func(t *testing.T) {

		body := http2.CreateUserRequest{
			ID:       userID.String(),
			Username: "userTest",
			Phone:    nil,
		}
		request, err := http.NewRequest(http.MethodPost, buildEndpointURL("/users"), buildRequest(t, body))
		require.NoError(t, err)
		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run(`Given a find user endpoint,
	when a request is sent,
	then it returns a valid response`, func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, buildEndpointURL(fmt.Sprintf("/users/%s", userID.String())), buildRequest(t, nil))
		require.NoError(t, err)
		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
