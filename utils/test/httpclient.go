package test

import (
	"net/http"

	"golang.org/x/net/context"
)

type MockHttpClientFactory struct {
}

func (f *MockHttpClientFactory) Client(ctx context.Context) *http.Client {
	return http.DefaultClient
}
