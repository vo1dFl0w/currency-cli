package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var RateURL = "https://v6.exchangerate-api.com/v6"

type CurrencyRateResponse struct {
	BaseCode        string `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func GetCurrencyRate(currencyCode string) (*CurrencyRateResponse, error) {
	apiKey := os.Getenv("EXCHANGERATE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("EXCHANGERATE_API_KEY not set")
	}

	url := fmt.Sprintf("%s/%s/latest/%s", RateURL, apiKey, currencyCode)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	var currencyRate CurrencyRateResponse
	if err := json.Unmarshal(data, &currencyRate); err != nil {
		return nil, fmt.Errorf("cannot parse response body: %w", err)
	}

	return &currencyRate, nil
}
