package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aubm/amount-converter-app/converter"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

type ConverterHandlers struct {
	Converter interface {
		RatesDef() *converter.RatesDef
		Convert(amount float64) map[string]float64
	} `inject:""`
	Logger interface {
		Infof(ctx context.Context, format string, args ...interface{})
	} `inject:""`
	Ctx ContextProviderInterface `inject:""`
}

func (h *ConverterHandlers) ConvertAmount(w http.ResponseWriter, r *http.Request) {
	ctx := h.Ctx.New(r)

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
	h.Logger.Infof(ctx, "Conversion done for amount %v: %v", amount, conversions)

	writeJSON(w, conversions, 200)
}

type FetchConvertAmountConfigurationAdapter struct {
	Converter interface {
		FetchConfiguration(ctx context.Context, url string) error
		RatesDef() *converter.RatesDef
	} `inject:""`
	Logger interface {
		Infof(ctx context.Context, format string, args ...interface{})
	} `inject:""`
	Ctx ContextProviderInterface `inject:""`
}

func (a *FetchConvertAmountConfigurationAdapter) Adapt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if a.Converter.RatesDef() == nil {
			if err := a.Converter.FetchConfiguration(ctx, API_URL); err != nil {
				log.Printf("Failed to init rates cache configuration: %v", err)
				writeError(w, SERVER_ERROR, 500)
				return
			}
			a.Logger.Infof(ctx, "Configuration loaded: %v", a.Converter.RatesDef())
		}
		next.ServeHTTP(w, r)
	})
}
