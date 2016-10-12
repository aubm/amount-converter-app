package test

import (
	"net/http"

	"golang.org/x/net/context"
)

type MockContextProvider struct{}

func (p MockContextProvider) New(r *http.Request) context.Context {
	return context.Background()
}
