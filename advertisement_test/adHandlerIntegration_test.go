package advertisement_test

import (
	"testing"
	"github.com/ankitggits/go-for-it/advertisement/handler"
	"net/http/httptest"
	"net/http"
	"log"
	"io/ioutil"
)

func TestFindAdByServiceHandler(t *testing.T) {
	requestHandler := handler.NewAdHandler().FindAdByServiceHandler
	responseString := integrationTest(requestHandler, "/service")
	printResponse(responseString)
	AssertNotNull(t, responseString, "Test FindAdByServiceHandler failed due to no response string")
}

func TestFindAdByCategoryHandler(t *testing.T) {
	requestHandler := handler.NewAdHandler().FindAdByCategoryHandler
	responseString := integrationTest(requestHandler, "/serice/IMR")
	printResponse(responseString)
	AssertNotNull(t, responseString, "Test FindAdByCategoryHandler failed due to no response string")
}

func TestSearchAdHandler(t *testing.T) {
	requestHandler := handler.NewAdHandler().SearchAdHandler
	responseString := integrationTest(requestHandler, "/serice/IMR/DFG_AS_1")
	printResponse(responseString)
	AssertNotNull(t, responseString, "Test SearchAdHandler failed due to no response string")
}

func printResponse(responseString string) {
	log.Printf("test response string : %s", responseString)
}

func integrationTest(requestHandler func(w http.ResponseWriter, r *http.Request), url string) string {

	ts := httptest.NewServer(http.HandlerFunc(requestHandler))

	defer ts.Close()

	res, err := http.Get(ts.URL+url)
	if err != nil {
		log.Fatal(err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(responseBytes)
}