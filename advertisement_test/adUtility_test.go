package advertisement_test

import (
	"github.com/ankitggits/go-for-it/advertisement/util"
	"testing"
)

func TestInit(t *testing.T) {
	store := util.Init("ads.json")

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

