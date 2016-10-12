package utils

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

type HttpClientFactory struct {
}

func (f *HttpClientFactory) Client(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}
