// Thin wraper around standard httptest library for easy mocking JSON HTTP responses.

package httptest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Endpoint represents mocked http endpoint.
type Endpoint struct {
	// Route is a URL path starting with slash.
	Route      string
	StatusCode int
	Content    string
}

// MockServer assebles mocked HTTP server with necessary routes.
func MockServer(endpoints ...*Endpoint) *httptest.Server {
	mux := http.NewServeMux()

	for _, endpoint := range endpoints {
		mux.HandleFunc(endpoint.Route, func(w http.ResponseWriter, request *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(endpoint.StatusCode)
			fmt.Fprint(w, endpoint.Content)
		})
	}

	return httptest.NewServer(mux)
}
