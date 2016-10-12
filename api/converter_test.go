package api

import (
	"net/http"
	"net/http/httptest"

	"github.com/aubm/amount-converter-app/converter/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConverterHandlers", func() {

	var mockConverterService *test.MockConverterService
	var converterHandlers ConverterHandlers
	var response *httptest.ResponseRecorder

	BeforeEach(func() {
		mockConverterService = &test.MockConverterService{}
		converterHandlers = ConverterHandlers{Converter: mockConverterService}
		response = httptest.NewRecorder()
	})

	Context("With a valid amount", func() {
		It("should convert the amount", func() {
			mockConverterService.On("Convert", float64(50)).Return(map[string]float64{"AUD": 11.11, "BGN": 22.22})
			request, _ := http.NewRequest("GET", "/?amount=50", nil)

			converterHandlers.ConvertAmount(response, request)

			Expect(response.Body.Bytes()).To(Equal([]byte(`{"AUD":11.11,"BGN":22.22}`)))
			Expect(response.Code).To(Equal(200))
		})
	})

	Context("With no amount", func() {
		It("should return an error", func() {
			request, _ := http.NewRequest("GET", "/", nil)

			converterHandlers.ConvertAmount(response, request)

			Expect(response.Body.Bytes()).To(Equal([]byte(`{"error":"Missing amount query parameter"}`)))
			Expect(response.Code).To(Equal(400))
		})
	})

	Context("With an invalid amount", func() {
		It("should return an error", func() {
			request, _ := http.NewRequest("GET", "/?amount=aaa", nil)

			converterHandlers.ConvertAmount(response, request)

			Expect(response.Body.Bytes()).To(Equal([]byte(`{"error":"Invalid amount parameter"}`)))
			Expect(response.Code).To(Equal(400))
		})
	})

})
