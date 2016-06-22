// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"regexp"
)

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"
)

var (
	// APIKey is the API key
	APIKey string
)

// SetAuth stores the credentials for later use
func SetAuth(key string) {
	APIKey = key
}

// HasAPIKey returns whether an API key is stored
func HasAPIKey() (string, bool) {
	if APIKey == "" {
		return "", false
	}
	return APIKey, true
}

// getPageNum returns the value of the page= parameter
func getPageNum(url string) (page string) {
	re := regexp.MustCompile(`page=(\d+)`)
	if m := re.FindStringSubmatch(url); len(m) >= 1 {
		return m[1]
	}
	return ""
}
