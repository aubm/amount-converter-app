package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aubm/amount-converter-app/converter"
)

type ConverterHandlers struct {
	Converter interface {
		FetchConfiguration(url string) error
		RatesDef() *converter.RatesDef
		Convert(amount float64) map[string]float64
	} `inject:""`
}

func (h *ConverterHandlers) ConvertAmount(w http.ResponseWriter, r *http.Request) {
	amountToConvert := r.URL.Query().Get("amount")
	if amountToConvert == "" {
		writeError(w, "Missing amount query parameter", 400)
		return
	}

	amount, err := strconv.ParseFloat(amountToConvert, 64)
	if err != nil {
		writeError(w, "Invalid amount parameter", 400)
		return
	}

	conversions := h.Converter.Convert(amount)
	writeJSON(w, conversions, 200)
}

type FetchConvertAmountConfigurationAdapter struct {
	Converter interface {
		FetchConfiguration(url string) error
		RatesDef() *converter.RatesDef
	} `inject:""`
}

func (a *FetchConvertAmountConfigurationAdapter) Adapt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.Converter.RatesDef() == nil {
			if err := a.Converter.FetchConfiguration(API_URL); err != nil {
				log.Printf("Failed to init rates cache configuration: %v", err)
				writeError(w, SERVER_ERROR, 500)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
