package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RatesDef struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func FetchConfiguration(url string) (*RatesDef, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch rates configuration url: %v", err)
	}

	rates := &RatesDef{}
	if err := json.NewDecoder(resp.Body).Decode(rates); err != nil {
		return nil, fmt.Errorf("Failed to parse rates configuration: %v", err)
	}

	return rates, nil
}

func Convert(amount float64, rates *RatesDef) map[string]float64 {
	conversions := map[string]float64{}
	for currency, rate := range rates.Rates {
		conversions[currency] = amount * rate
	}
	return conversions
}
