package utils

import "os"

// GetBaseURL returns baseURL of request
func GetBaseURL() string {
	var baseURL string
	env := os.Getenv("GO_ENV")
	if env == "development" {
		baseURL = "/"
	} else {
		baseURL = "/api/"
	}
	return baseURL
}
