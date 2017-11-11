// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"
)

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

// AddQueryParameters adds query parameters to the URL.
func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// prepareRequest insert all pre-defined stuff
func prepareRequest(method, what string, opts map[string]string) (req *http.Request) {
	endPoint := apiEndpoint + fmt.Sprintf("/%s/", what)
	key, ok := HasAPIKey()

	// Insert key
	if ok {
		opts["key"] = key
	}

	baseURL := AddQueryParameters(endPoint, opts)

	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		log.Printf("error parsing %s: %v", baseURL, err)
		return &http.Request{}
	}

	myurl, err := url.Parse(baseURL)

	req.Header.Set("Host", myurl.Host)
	req.Header.Set("User-Agent", fmt.Sprintf("ripe-atlas/%s", ourVersion))

	return
}

// handleAPIResponse check status code & errors from the API
func handleAPIResponse(r *http.Response) (err error) {
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

		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		err = json.Unmarshal(body, &aerr)
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

	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, &aerr)
	if err != nil {
		log.Printf("Error handling error: %s - %v", r.Body, err)
	}

	err = fmt.Errorf("status: %d code: %d - r:%v",
		aerr.Error.Status,
		aerr.Error.Code,
		aerr.Error.Detail)
	return
}
