package handler

import (
	"net/http"
	"regexp"
)

// Route type contains regular expressions of URLs and their respective handlers
type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler is a custom implements of http Handler and contains all the route information for routing
// It solves the purpose of identifying Restful URLs with path parameters
type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handle(pattern string, handler http.Handler) {
	h.routes = append(h.routes, &route{regexp.MustCompile(pattern), handler})
}

// Overrides http handler's HandleFunc and provides implementation for executing handlers specific to patterns registered
func (h *RegexpHandler) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.routes = append(h.routes, &route{regexp.MustCompile(pattern), http.HandlerFunc(handler)})
}

// serves http requests based on the registered routes.
// if not matches with existing registry , responds with 404 status code
func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
