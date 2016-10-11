package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aubm/amount-converter-app/converter"
)

const (
	API_URL      = "https://api.fixer.io/latest"
	SERVER_ERROR = "An error occured"
	ADDR         = ":8080"
)

func init() {
	http.HandleFunc("/convert", convertAmount)
}

func main() {
	fmt.Printf("Server started on %v\n", ADDR)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		panic(err)
	}
}

var ratesCache *converter.RatesDef

func convertAmount(w http.ResponseWriter, r *http.Request) {
	if ratesCache == nil {
		var err error
		ratesCache, err = converter.FetchConfiguration(API_URL)
		if err != nil {
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

	conversions := converter.Convert(amount, ratesCache)
	writeJSON(w, conversions, 200)
}

func writeError(w http.ResponseWriter, error string, statusCode int) {
	writeJSON(w, map[string]string{
		"error": error,
	}, statusCode)
}

func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	w.Write(b)
}
