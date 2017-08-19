package model

type Ad struct {
	AdKey      string `json:"ad_key"`
	AdProvider string `json:"ad_provider"`
	AdText     string `json:"ad_text"`
}
