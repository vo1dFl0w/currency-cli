package commands

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/currency-cli/internal/api/exchange"
)

func TestConvert_ValidJSON(t *testing.T) {
	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"base_code": "USD",
			"target_code": "RUB",
			"conversion_rate": 78.4447
		}`)
	})

	resp, err := exchange.Convert("USD", "RUB")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.BaseCode != "USD" || resp.TargetCode != "RUB" || resp.ConversionRate != 78.4447 {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestConvert_NoAPIKey(t *testing.T) {
	os.Unsetenv("EXCHANGERATE_API_KEY")

	_, err := exchange.Convert("USD", "RUB")
	if err == nil || !strings.Contains(err.Error(), "EXCHANGERATE_API_KEY not set") {
		t.Fatalf("expected API key error, got: %v", err)
	}
}

func TestConvert_HTTPRequestError(t *testing.T) {
	os.Setenv("EXCHANGERATE_API_KEY", "fakeapikey")
	exchange.ConvertURL = "http://localhost:0"

	_, err := exchange.Convert("USD", "RUB")
	if err == nil || !strings.Contains(err.Error(), "request error") {
		t.Fatalf("expected request error, got: %v", err)
	}
}

func TestConvert_HTTPStatusError(t *testing.T) {
	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, err := exchange.Convert("USD", "RUB")
	if err == nil || !strings.Contains(err.Error(), "response error") {
		t.Fatalf("expected response error, got: %v", err)
	}
}

func TestConvert_InvalidJSON(t *testing.T) {
	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"base_code": "USD",	 
		}`)
	})

	_, err := exchange.Convert("USD", "RUB")
	if err == nil || !strings.Contains(err.Error(), "cannot parse response body") {
		t.Fatalf("expected JSON parse error, got: %v", err)
	}
}

func TestConvert_EmptyResponse(t *testing.T) {
	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	_, err := exchange.Convert("USD", "RUB")
	if err == nil || !strings.Contains(err.Error(), "cannot parse response body") {
		t.Fatalf("expected parse error for response body, got: %v", err)
	}
}

func TestConvert_RequestPath(t *testing.T) {
	var path string

	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		fmt.Fprint(w, `{"base_code": "USD", "target_code": "RUB", "conversion_rate": 78.4447}`)
	})

	_, _ = exchange.Convert("USD", "RUB")

	expected := "/fakeapikey/pair/USD/RUB"
	if !strings.HasSuffix(path, expected) {
		t.Errorf("unexpected request path: got %s, want suffix %s", path, expected)
	}
}

func setupConvertTestServer(t *testing.T, h http.HandlerFunc) *httptest.Server {
	ts := httptest.NewServer(h)
	t.Cleanup(func() {
		ts.Close()
	})

	originalURL := exchange.ConvertURL
	exchange.ConvertURL = ts.URL
	t.Cleanup(func() {
		exchange.ConvertURL = originalURL
	})

	if err := os.Setenv("EXCHANGERATE_API_KEY", "fakeapikey"); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	return ts
}
