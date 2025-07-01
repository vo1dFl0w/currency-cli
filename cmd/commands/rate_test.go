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

func TestRate_ValidJSON(t *testing.T) {
	setupRateTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"base_code": "USD",
			"conversion_rates": {
				"RUB": 78.4447
			}
		}`)
	})

	resp, err := exchange.GetCurrencyRate("USD")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.BaseCode != "USD" || resp.ConversionRates["RUB"] != 78.4447 {
		t.Fatalf("unexpected response: %v", err)
	}
}

func TestRate_InvalidJSON(t *testing.T) {
	setupRateTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"base_code": "USD",
		}`)
	})

	_, err := exchange.GetCurrencyRate("USD")
	if err == nil || !strings.Contains(err.Error(), "cannot parse response body") {
		t.Fatalf("expected JSON parse error, got: %v", err)
	}
}

func TestRate_NoAPIKey(t *testing.T) {
	os.Unsetenv("EXCHANGERATE_API_KEY")

	_, err := exchange.GetCurrencyRate("USD")
	if err == nil || !strings.Contains(err.Error(), "EXCHANGERATE_API_KEY not set") {
		t.Fatalf("expected API key error, got: %v", err)
	}
}

func TestRate_HTTPRequestError(t *testing.T) {
	os.Setenv("EXCHANGERATE_API_KEY", "fakeapikey")
	exchange.RateURL = "http://localhost:0"

	_, err := exchange.GetCurrencyRate("USD")
	if err == nil || !strings.Contains(err.Error(), "request error") {
		t.Fatalf("expected request error, got: %v", err)
	}
}

func TestRate_HTTPStatusError(t *testing.T) {
	setupConvertTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	})

	_, err := exchange.GetCurrencyRate("USD")
	if err == nil || !strings.Contains(err.Error(), "response error") {
		t.Fatalf("expected response error, got: %v", err)
	}
}

func TestRate_EmptyResponse(t *testing.T) {
	setupRateTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	_, err := exchange.GetCurrencyRate("USD")
	if err == nil || !strings.Contains(err.Error(), "cannot parse response body") {
		t.Fatalf("expected parse error for response body, got: %v", err)
	}
}

func setupRateTestServer(t *testing.T, h http.HandlerFunc) *httptest.Server {
	ts := httptest.NewServer(h)
	t.Cleanup(func() {
		ts.Close()
	})

	originalURL := exchange.RateURL
	exchange.RateURL = ts.URL
	t.Cleanup(func() {
		exchange.RateURL = originalURL
	})

	if err := os.Setenv("EXCHANGERATE_API_KEY", "fakeapikey"); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	return ts
}
