// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"log"
	"regexp"
)

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"

	ourVersion = "0.12"
)

// HasAPIKey returns whether an API key is stored
func HasAPIKey() (string, bool) {
	if ctx.config.APIKey == "" {
		return "", false
	}
	return ctx.config.APIKey, true
}

// GetVersion returns the API wrapper version
func GetVersion() string {
	return ourVersion
}

// getPageNum returns the value of the page= parameter
func getPageNum(url string) (page string) {
	re := regexp.MustCompile(`page=(\d+)`)
	if m := re.FindStringSubmatch(url); len(m) >= 1 {
		return m[1]
	}
	return ""
}

// prepareRequest insert all pre-defined stuff
func prepareRequest(what string) (req rest.Request) {
	endPoint := apiEndpoint + fmt.Sprintf("/%s/", what)
	key, ok := HasAPIKey()

	// Add at least one option, the APIkey if present
	hdrs := make(map[string]string)
	opts := make(map[string]string)

	// Insert our sig
	hdrs["User-Agent"] = fmt.Sprintf("ripe-atlas/%s", ourVersion)

	// Insert key
	if ok {
		opts["key"] = key
	}

	req = rest.Request{
		BaseURL:     endPoint,
		Headers:     hdrs,
		QueryParams: opts,
	}
	return
}

// handleAPIResponse check status code & errors from the API
func handleAPIResponse(r *rest.Response) (err error) {
	if r == nil {
		return fmt.Errorf("Error: r is nil!")
	}

	// Everything is fine
	if r.StatusCode == 0 {
		return nil
	}

	// Everything is fine too
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}

	// Check this condition
	if r.StatusCode >= 300 && r.StatusCode <= 399 {
		var aerr APIError

		err = json.Unmarshal([]byte(r.Body), &aerr)
		if err != nil {
			log.Printf("Error handling error: %s - %v", r.Body, err)
		}

		log.Printf("Info 3XX status: %d code: %d - r:%v\n",
			aerr.Error.Status,
			aerr.Error.Code,
			aerr.Error.Detail)
		return nil
	}

	// EVerything else is an error
	var aerr APIError

	err = json.Unmarshal([]byte(r.Body), &aerr)
	if err != nil {
		log.Printf("Error handling error: %s - %v", r.Body, err)
	}

	err = fmt.Errorf("status: %d code: %d - r:%v",
		aerr.Error.Status,
		aerr.Error.Code,
		aerr.Error.Detail)
	return
}
