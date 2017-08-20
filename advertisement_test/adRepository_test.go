package advertisement_test

import (
	"testing"
	"github.com/ankitggits/go-for-it/advertisement/repo"
)

func TestRandomAd(t *testing.T) {
	repo := repo.NewAdRepository()
	found, _ := repo.FindRandomAd()
	if !found {
		t.Errorf("Test Random Ad failed , Expected record but found none")
	}
}

func TestFindRandomAdByCategory(t *testing.T) {
	repo := repo.NewAdRepository()
	found, ad := repo.FindRandomAdByCategory("IMR")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}else if !(ad.AdKey=="DFG_AS_1" || ad.AdKey=="TFG_AS_1" || ad.AdKey=="TFG_AS_2"){
		t.Errorf("Test Random Ad by category failed , found other category record")
	}
}

func TestFindAdUnknownCategory(t *testing.T) {
	repo := repo.NewAdRepository()
	found, _ := repo.FindRandomAdByCategory("ABC")
	if found {
		t.Errorf("Test Random Ad by category failed , Expected none but found one")
	}
}

func TestFindAdByAdCategoryAndKey(t *testing.T) {
	repo := repo.NewAdRepository()
	_, category := repo.FindAdCategory("IMR")
	found, _ := repo.FindAdByAdCategoryAndKey(category, "DFG_AS_1")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}
}

func TestFindAdByAdCategoryAndProvider(t *testing.T) {
	repo := repo.NewAdRepository()
	_, category := repo.FindAdCategory("IMR")
	found, _ := repo.FindAdByAdCategoryAndProvider(category, "The_Future_Group_AS")
	if !found {
		t.Errorf("Test Random Ad by category failed , Expected record but found none")
	}
}