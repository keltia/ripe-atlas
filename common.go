// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
    "regexp"
    "errors"
)

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"

	ourVersion = "0.8"
)

// APIKey is the API key
var APIKey string

// ErrInvalidMeasurementType is a new error
var ErrInvalidMeasurementType = errors.New("invalid measurement type")

// ErrInvalidAPIKey is returned when the key is invalid
var ErrInvalidAPIKey = errors.New("invalid API key")

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
