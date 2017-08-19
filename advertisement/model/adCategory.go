package model

type AdCategory struct {
	AdCategory string   `json:"ad_category"`
	Ads        []Ad 	`json:"ads"`
}
