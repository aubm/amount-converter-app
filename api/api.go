package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

const (
	API_URL      = "https://api.fixer.io/latest"
	SERVER_ERROR = "An error occured"
)

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

type ContextProviderInterface interface {
	New(r *http.Request) context.Context
}

type ContextProvider struct{}

func (p ContextProvider) New(r *http.Request) context.Context {
	return appengine.NewContext(r)
}
