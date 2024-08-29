//go:build functional

package functional

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	ApiHost    = "http://localhost:8083"
	httpClient *http.Client
)

func TestMain(m *testing.M) {
	buildHttpClient()
	code := m.Run()
	os.Exit(code)
}

func buildEndpointURL(uri string) string {
	return fmt.Sprintf("%s%s", ApiHost, uri)
}

func buildHttpClient() {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 5 * time.Second}
	}
}
