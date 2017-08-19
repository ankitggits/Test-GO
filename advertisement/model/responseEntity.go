package model

import "time"

type responseEntity struct {
	Response  				Ad 				`json:"response,omitempty"`
	RequestPath  			string 			`json:"requestPath,omitempty"`
	RequestExecutionTime  	string			`json:"executionTime"`
}

// A response constructor to construct http json response, will contain requested path and execution time
// along with the Advertisement
func NewResponseEntity(ad Ad, requestPath string, timeTaken time.Duration) *responseEntity {
	return &responseEntity{ad, requestPath, timeTaken.String()}
}
