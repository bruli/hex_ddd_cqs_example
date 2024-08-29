//go:build functional

package functional

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func buildRequest(t *testing.T, requestBody any) io.Reader {
	var buff bytes.Buffer
	body, err := json.Marshal(requestBody)
	require.NoError(t, err, "failed to marshal request body")
	buff.Write(body)

	return &buff
}
