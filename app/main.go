package main

import (
	"fmt"
	"net/http"

	"github.com/aubm/amount-converter-app/api"
	"github.com/aubm/amount-converter-app/converter"
	"github.com/facebookgo/inject"
)

const ADDR = ":8080"

func init() {
	converterHandlers := &api.ConverterHandlers{}
	converterService := &converter.ConverterService{}

	if err := inject.Populate(converterHandlers, converterService); err != nil {
		panic(err)
	}

	http.HandleFunc("/convert", converterHandlers.ConvertAmount)
}

func main() {
	fmt.Printf("Server started on %v\n", ADDR)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		panic(err)
	}
}
