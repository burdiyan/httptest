// Thin wrapper around standard httptest library for easy mocking JSON HTTP responses.

package httptest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// Server is an alias of the server from httptest.
type Server struct {
	*httptest.Server
}

// Endpoint represents mocked http endpoint.
type Endpoint struct {
	// Route is a URL path starting with slash.
	Route       string
	StatusCode  int
	Content     string
	ContentType string
}

// JSONEndpoint creates the endpoint with content type for JSON.
func JSONEndpoint(route string, statusCode int, content string) *Endpoint {
	return &Endpoint{
		Route:       route,
		StatusCode:  statusCode,
		Content:     content,
		ContentType: "application/json",
	}
}

// HTMLEndpoint created the endpoint with content type for HTML.
func HTMLEndpoint(route string, statusCode int, content string) *Endpoint {
	return &Endpoint{
		Route:       route,
		StatusCode:  statusCode,
		Content:     content,
		ContentType: "text/html",
	}
}

// MockServer assembles mocked HTTP server with necessary routes.
func MockServer(endpoints ...*Endpoint) *Server {
	mux := http.NewServeMux()

	for _, e := range endpoints {
		mux.HandleFunc(e.Route, func(w http.ResponseWriter, request *http.Request) {
			w.Header().Add("Content-Type", e.ContentType)
			w.WriteHeader(e.StatusCode)
			fmt.Fprint(w, e.Content)
		})
	}

	return &Server{httptest.NewServer(mux)}
}
