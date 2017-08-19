package main

import (
	"net/http"
	"log"
	"time"
	"github.com/ankitggits/go-for-it/advertisement/repo"
	"github.com/ankitggits/go-for-it/advertisement/handler"
)

// Trace requests and their execution timing
func traceableHandler(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		t1:= time.Now()
		next.ServeHTTP(w,r)
		t2:= time.Now()
		log.Printf("[%s] %q %v\n" , r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

// Handler adds Response header, For this example only adds content type as json
func headersHandler(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)
}

// Handler validates Request Method, For this example only allows Get Request Method
func requestValidatorHandler(next http.Handler) http.Handler{
	fn:= func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)
}

// Application startup with InMem storage initialization and routing.
// Uses Custom RegexpHandler for evaluating and validating Restful URL based Path params along with
// adds chains of handlers for request method validation, logging trace and adding response headers
func main() {
	repo.Init("ads.json")
	regexHnd := new(handler.RegexpHandler)
	regexHnd.HandleFunc("/service$",  handler.FindAdByServiceHandler)
	regexHnd.HandleFunc("/service/[a-zA-Z_0-9]*$", handler.FindAdByCategoryHandler)
	regexHnd.HandleFunc("/service/[a-zA-Z_0-9]*/[a-zA-Z._0-9]*$", handler.SearchAdHandler)
	log.Fatal(http.ListenAndServe(":8090", requestValidatorHandler(traceableHandler(headersHandler(regexHnd)))))
}
