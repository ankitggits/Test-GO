package handler

import (
	"net/http"
	"time"
	"github.com/ankitggits/go-for-it/advertisement/util"
	"github.com/ankitggits/go-for-it/advertisement/repo"
)

// Api will find a random ad from any category , in case of unavailability return 404 http status code
func FindAdByServiceHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	found, ad := repo.FindRandomAd()
	if found {
		util.WriteJson(w, r, ad, startTime)
	}else{
		http.NotFound(w,r)
	}
}

// Api will find a random ad from given category in URL path param , in case of unavailability return 404 http status code
func FindAdByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	found, ad := repo.FindRandomAdByCategory(util.GetPathParam(r.URL.Path, 0))
	if found {
		util.WriteJson(w, r, ad, startTime)
	}else{
		http.NotFound(w,r)
	}
}

// Api will find a random ad from given category and search text in URL path param , in case of unavailability return 404 http status code
// Note: search text can contain adKey or adProvider. adKey has given priority over provider
func SearchAdHandler(w http.ResponseWriter, r *http.Request) {
	startTime:=time.Now()
	category := util.GetPathParam(r.URL.Path, 1)
	searchText := util.GetPathParam(r.URL.Path, 0)
	adFound, adCategory := repo.FindAdCategory(category)
	if !adFound{
		http.NotFound(w,r)
	}
	keyFound, keyAd := repo.FindAdByAdCategoryAndKey(adCategory, searchText)
	if keyFound {
		util.WriteJson(w, r, keyAd, startTime)
	}else {
		providerFound, providerAd := repo.FindAdByAdCategoryAndProvider(adCategory, searchText)
		if providerFound {
			util.WriteJson(w, r, providerAd, startTime)
		}else {
			http.NotFound(w,r)
		}
	}
}


