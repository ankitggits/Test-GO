package handler

import (
	"net/http"
	"time"
	"log"
)

type ChainHandler interface {
	Next(next http.Handler) http.Handler
}

type traceableHandler struct {}

func NewTraceableHandler() ChainHandler{
	return &traceableHandler{}
}

// Trace requests and their execution timing
func (th traceableHandler) Next(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		t1:= time.Now()
		next.ServeHTTP(w,r)
		t2:= time.Now()
		log.Printf("[%s] %q %v\n" , r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

type responseHeaderHandler struct {}

func NewResponseHeaderHandler() ChainHandler{
	return &responseHeaderHandler{}
}

// Handler adds Response header, For this example only adds content type as json
func (rhh responseHeaderHandler) Next(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)
}

type requestValidatorHandler struct {}

func NewRequestValidatorHandler() ChainHandler{
	return &requestValidatorHandler{}
}

// Handler validates Request Method, For this example only allows Get Request Method
func (rvh requestValidatorHandler) Next(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)
}
