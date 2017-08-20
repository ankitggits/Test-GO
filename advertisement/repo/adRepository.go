package repo

import (
	"math/rand"
	"github.com/ankitggits/go-for-it/advertisement/model"
	"github.com/ankitggits/go-for-it/advertisement/util"
	"github.com/ankitggits/go-for-it/advertisement/config"
)

type SuperRepository interface {
	GetStore() *model.AdStore
	FindRandomAd() (bool, model.Ad)
	FindRandomAdByCategory(cat string) (bool,model.Ad)
	FindAdCategory(category string) (bool,model.AdCategory)
	FindAdByAdCategoryAndKey(category model.AdCategory, key string) (bool,model.Ad)
	FindAdByAdCategoryAndProvider(category model.AdCategory, provider string) (bool,model.Ad)
}

type adRepository struct {
	store *model.AdStore
}

func NewAdRepository() SuperRepository {
	return &adRepository{util.Init(config.FILE_PATH)}
}

// for testing purpose only
func (repo adRepository) GetStore() *model.AdStore{
	return repo.store
}

// Find Random ad , In case of unavailability return false
func (repo adRepository) FindRandomAd() (bool, model.Ad) {
	if len(repo.store.AdCategories)>0{
		numberOfCategories:=len(repo.store.AdCategories)
		randomCategoryIndex := rand.Intn(numberOfCategories)
		return repo.FindRandomAdByCategory(repo.store.AdCategories[randomCategoryIndex].AdCategory)
	}
	return false, model.Ad{}
}

// Find Random ad by given category, In case of unavailability return false
func (repo adRepository) FindRandomAdByCategory(cat string) (bool,model.Ad){
	for i := 0; i < len(repo.store.AdCategories); i++ {
		if repo.store.AdCategories[i].AdCategory==cat && len(repo.store.AdCategories[i].Ads)>0{
			numberOfAds := len(repo.store.AdCategories[i].Ads)
			randomAdIndex := rand.Intn(numberOfAds)
			return true,repo.store.AdCategories[i].Ads[randomAdIndex]
		}
	}
	return false,model.Ad{}
}

// Find Ad Category by given name , In case of unavailability return false
func (repo adRepository) FindAdCategory(category string) (bool,model.AdCategory){
	for i :=0; i < len(repo.store.AdCategories); i++ {
		adCategory := repo.store.AdCategories[i]
		if adCategory.AdCategory==category{
			return true, adCategory
		}
	}
	return false,model.AdCategory{}
}

// Find Ad within given category by adkey , In case of unavailability return false
func (repo adRepository) FindAdByAdCategoryAndKey(category model.AdCategory, key string) (bool,model.Ad){
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
func (repo adRepository) FindAdByAdCategoryAndProvider(category model.AdCategory, provider string) (bool,model.Ad){
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