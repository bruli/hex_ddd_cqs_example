//go:build functional

package functional

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHomepage(t *testing.T) {
	t.Run(`Given a homepage endpoint,
	when a request is sent,
	then it returns an OK status`, func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, buildEndpointURL("/"), nil)
		require.NoError(t, err)
		resp, err := httpClient.Do(request)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
