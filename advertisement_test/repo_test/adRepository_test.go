package repo_test

import (
	"testing"
	"github.com/ankitggits/go-for-it/advertisement/repo"
)

func TestInit(t *testing.T) {
	repo.Init("ads_test.json")
	store := repo.GetStore()
	expectedCategories := 2
	//test number of categories initialized
	if len(store.AdCategories)!= expectedCategories {
		t.Errorf("Test Init failed , expected store initialization with %d categories but found %d", expectedCategories, len(store.AdCategories))
	}
	expectedAdsOfIMR := 3
	//test number of ads initialized for IMR category
	if len(store.AdCategories[0].Ads)!= expectedAdsOfIMR {
		t.Errorf("Test Init failed , expected store initialization with %d categories but found %d", expectedAdsOfIMR, len(store.AdCategories[0].Ads))
	}

	expectedAdsOfGames := 3
	//test number of ads initialized for Games category
	if len(store.AdCategories[1].Ads)!= expectedAdsOfGames {
		t.Errorf("Test Init failed , expected store initialization with %d categories but found %d", expectedAdsOfGames, len(store.AdCategories[1].Ads))
	}
}

func TestRandomAd(t *testing.T) {
	repo.Init("ads_test.json")
	found, _ := repo.FindRandomAd()
	if !found {
		t.Errorf("Test Random Ad failed , Expected record but found none")
	}
}

func TestFindRandomAdByCategory(t *testing.T) {
	repo.Init("ads_test.json")
	found, ad := repo.FindRandomAdByCategory("IMR")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}else if !(ad.AdKey=="DFG_AS_1" || ad.AdKey=="TFG_AS_1"){
		t.Errorf("Test Random Ad by category failed , found other category record")
	}
}

func TestFindAdUnknownCategory(t *testing.T) {
	repo.Init("ads_test.json")
	found, _ := repo.FindRandomAdByCategory("ABC")
	if found {
		t.Errorf("Test Random Ad by category failed , Expected none but found one")
	}
}

func TestFindAdByAdCategoryAndKey(t *testing.T) {
	repo.Init("ads_test.json")
	_, category := repo.FindAdCategory("IMR")
	found, _ := repo.FindAdByAdCategoryAndKey(category, "DFG_AS_1")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}
}

func TestFindAdByAdCategoryAndProvider(t *testing.T) {
	repo.Init("ads_test.json")
	_, category := repo.FindAdCategory("IMR")
	found, _ := repo.FindAdByAdCategoryAndProvider(category, "The_Future_Group_AS")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}
}