package main

import (
	"net/http"
	"log"
	"github.com/ankitggits/go-for-it/advertisement/handler"
)

// Application startup with InMem storage initialization and routing.
// Uses Custom RegexpHandler for evaluating and validating Restful URL based Path params along with
// adds chains of handlers for request method validation, logging trace and adding response headers
func main() {
	//Advertisement Handler
	adHandler := handler.NewAdHandler()
	//Regular Expression Handler for matching the pattern of restful type and fetching path params
	regexHnd  := new(handler.RegexpHandler)

	//Chain Handlers
	requestValidatorHandler := handler.NewRequestValidatorHandler()
	traceableHandler := handler.NewTraceableHandler()
	headersHandler := handler.NewResponseHeaderHandler()

	//Handler Chain configuration
	handlerChain := requestValidatorHandler.Next(
							traceableHandler.Next(
							headersHandler.Next(regexHnd)))

	regexHnd.HandleFunc("/service$",  adHandler.FindAdByServiceHandler)
	regexHnd.HandleFunc("/service/[a-zA-Z_0-9]*$", adHandler.FindAdByCategoryHandler)
	regexHnd.HandleFunc("/service/[a-zA-Z_0-9]*/[a-zA-Z._0-9]*$", adHandler.SearchAdHandler)
	log.Fatal(http.ListenAndServe(":8090", handlerChain))
}
