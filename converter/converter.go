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

type ConverterService struct {
	ratesDef *RatesDef
}

func (c *ConverterService) RatesDef() *RatesDef {
	return c.ratesDef
}

func (c *ConverterService) FetchConfiguration(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to fetch rates configuration url: %v", err)
	}

	c.ratesDef = &RatesDef{}
	if err := json.NewDecoder(resp.Body).Decode(c.ratesDef); err != nil {
		return fmt.Errorf("Failed to parse rates configuration: %v", err)
	}

	return nil
}

func (c *ConverterService) Convert(amount float64) map[string]float64 {
	conversions := map[string]float64{}
	for currency, rate := range c.ratesDef.Rates {
		conversions[currency] = amount * rate
	}
	return conversions
}
