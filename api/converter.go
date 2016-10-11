package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aubm/amount-converter-app/converter"
)

type ConverterHandlers struct {
	Converter *converter.ConverterService `inject:""`
}

func (h *ConverterHandlers) ConvertAmount(w http.ResponseWriter, r *http.Request) {
	if h.Converter.RatesDef == nil {
		if err := h.Converter.FetchConfiguration(API_URL); err != nil {
			log.Printf("Failed to init rates cache configuration: %v", err)
			writeError(w, SERVER_ERROR, 500)
			return
		}
	}

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
