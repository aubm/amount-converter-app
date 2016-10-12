package converter

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	utilsTest "github.com/aubm/amount-converter-app/utils/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("ConverterService", func() {

	var converterService ConverterService
	var server *httptest.Server

	BeforeEach(func() {
		converterService = ConverterService{
			HTTP: &utilsTest.MockHttpClientFactory{},
			ratesDef: &RatesDef{
				Rates: map[string]float64{
					"AUD": 1.4679,
					"BGN": 1.9558,
					"BRL": 3.5634,
					"CAD": 1.4653,
					"CHF": 1.0938,
				},
			},
		}
	})

	BeforeEach(func() {
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{ "base": "EUR", "date":"2016-10-11",
			  "rates": { "AUD": 1.4679, "BGN": 1.9558, "BRL": 3.5634, "CAD": 1.4653 } }`)
		}))
	})

	AfterEach(func() {
		server.Close()
	})

	It("should fetch the configuration from the server", func() {
		ctx := context.Background()

		converterService.FetchConfiguration(ctx, server.URL)

		Expect(converterService.RatesDef()).To(Equal(&RatesDef{
			Base: "EUR",
			Date: "2016-10-11",
			Rates: map[string]float64{
				"AUD": 1.4679,
				"BGN": 1.9558,
				"BRL": 3.5634,
				"CAD": 1.4653,
			},
		}))
	})

	It("should convert the given amount", func() {
		amounts := converterService.Convert(50)

		Expect(amounts).To(Equal(
			map[string]float64{
				"AUD": 73.395,
				"BGN": 97.78999999999999,
				"BRL": 178.17000000000002,
				"CAD": 73.265,
				"CHF": 54.690000000000005,
			},
		))
	})

})
