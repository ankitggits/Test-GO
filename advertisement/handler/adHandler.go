package handler

import (
	"net/http"
	"time"
	"github.com/ankitggits/go-for-it/advertisement/util"
	"github.com/ankitggits/go-for-it/advertisement/repo"
)

type SuperHandler interface {
	FindAdByServiceHandler(w http.ResponseWriter, r *http.Request)
	FindAdByCategoryHandler(w http.ResponseWriter, r *http.Request)
	SearchAdHandler(w http.ResponseWriter, r *http.Request)
}

type adHandler struct {
	repo repo.SuperRepository
}

func NewAdHandler() SuperHandler{
	return &adHandler{repo.NewAdRepository()}
}

// Api will find a random ad from any category , in case of unavailability return 404 http status code
func (handler adHandler) FindAdByServiceHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	found, ad := handler.repo.FindRandomAd()
	if found {
		util.WriteJson(w, r, ad, startTime)
	}else{
		http.NotFound(w,r)
	}
}

// Api will find a random ad from given category in URL path param , in case of unavailability return 404 http status code
func (handler adHandler) FindAdByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	found, ad := handler.repo.FindRandomAdByCategory(util.GetPathParam(r.URL.Path, 0))
	if found {
		util.WriteJson(w, r, ad, startTime)
	}else{
		http.NotFound(w,r)
	}
}

// Api will find a random ad from given category and search text in URL path param , in case of unavailability return 404 http status code
// Note: search text can contain adKey or adProvider. adKey has given priority over provider
func (handler adHandler) SearchAdHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	category := util.GetPathParam(r.URL.Path, 1)
	searchText := util.GetPathParam(r.URL.Path, 0)
	adFound, adCategory := handler.repo.FindAdCategory(category)
	if !adFound{
		http.NotFound(w,r)
	}
	keyFound, keyAd := handler.repo.FindAdByAdCategoryAndKey(adCategory, searchText)
	if keyFound {
		util.WriteJson(w, r, keyAd, startTime)
	}else {
		providerFound, providerAd := handler.repo.FindAdByAdCategoryAndProvider(adCategory, searchText)
		if providerFound {
			util.WriteJson(w, r, providerAd, startTime)
		}else {
			http.NotFound(w,r)
		}
	}
}


