package util

import (
	"net/http"
	"time"
	"encoding/json"
	"strings"
	"github.com/ankitggits/go-for-it/advertisement/model"
	"io/ioutil"
	"log"
	"os"
)

// Initialize in memory data storage, Responsible for unmarshal and and initialization of AdStore
func Init(path string) *model.AdStore {
	adStore := new(model.AdStore)
	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, adStore)
	log.Printf("DB initialized with %d categories", len(adStore.AdCategories))
	return adStore
}

// Utility function features creating new Response Entity and writing it to json Response
func WriteJson(w http.ResponseWriter, r *http.Request, advertisement model.Ad, startTime time.Time ){
	jData, err := json.Marshal(model.NewResponseEntity(advertisement, r.URL.String(), time.Now().Sub(startTime)))
	if err != nil {
		panic(err)
		return
	}
	w.Write(jData)
}

// Utility function features fetching path param by splitting URl by '/' and finding value by given index from last
// Chances of errors are less because of having regex pattern based routing
func GetPathParam(url string, indexFromLast int) (string) {
	p := strings.Split(url, "/")
	return p[len(p)-1-indexFromLast]
}