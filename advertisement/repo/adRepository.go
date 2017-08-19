package repo

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"math/rand"
	"log"
	"github.com/ankitggits/go-for-it/advertisement/model"
)

var store model.AdStore

// Initialize in memory data storage, Responsible for unmarshal and and initialization of AdStore
func Init(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &store)
	log.Printf("DB initialized with %d categories", len(store.AdCategories))
}

// for testing purpose only
func GetStore() model.AdStore{
	return store
}

// Find Random ad , In case of unavailability return false
func FindRandomAd() (bool, model.Ad) {
	if len(store.AdCategories)>0{
		numberOfCategories:=len(store.AdCategories)
		randomCategoryIndex := rand.Intn(numberOfCategories)
		return FindRandomAdByCategory(store.AdCategories[randomCategoryIndex].AdCategory)
	}
	return false, model.Ad{}
}

// Find Random ad by given category, In case of unavailability return false
func FindRandomAdByCategory(cat string) (bool,model.Ad){
	for i := 0; i < len(store.AdCategories); i++ {
		if store.AdCategories[i].AdCategory==cat && len(store.AdCategories[i].Ads)>0{
			numberOfAds := len(store.AdCategories[i].Ads)
			randomAdIndex := rand.Intn(numberOfAds)
			return true,store.AdCategories[i].Ads[randomAdIndex]
		}
	}
	return false,model.Ad{}
}

// Find Ad Category by given name , In case of unavailability return false
func FindAdCategory(category string) (bool,model.AdCategory){
	for i :=0; i < len(store.AdCategories); i++ {
		adCategory := store.AdCategories[i]
		if adCategory.AdCategory==category{
			return true, adCategory
		}
	}
	return false,model.AdCategory{}
}

// Find Ad within given category by adkey , In case of unavailability return false
func FindAdByAdCategoryAndKey(category model.AdCategory, key string) (bool,model.Ad){
	for j :=0; j < len(category.Ads); j++ {
		ad := category.Ads[j]
		if ad.AdKey==key{
			return true,ad
		}
	}
	return false,model.Ad{}
}

// Find Ad within given category by adProvider. When multiple ads found for same provider returns random Ad
// In case of unavailability return false
func FindAdByAdCategoryAndProvider(category model.AdCategory, provider string) (bool,model.Ad){
	var ads []model.Ad
	for i :=0; i < len(category.Ads); i++ {
		ad := category.Ads[i]
		if ad.AdProvider==provider{
			ads = append(ads, ad)
		}
	}
	if len(ads)>0 {
		numberOfProviderAds := len(ads)
		randomAdIndex := rand.Intn(numberOfProviderAds)
		return true, ads[randomAdIndex]
	}
	return false,model.Ad{}
}