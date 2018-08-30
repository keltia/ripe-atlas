// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"
)

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
	if len(queryParams) == 0 {
		return baseURL
	}
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// addAPIKey insert the key into options if needed
func (c *Client) addAPIKey(opts map[string]string) map[string]string {
	key, ok := c.HasAPIKey()
	// Insert key
	if ok {
		opts["key"] = key
	}
	return opts
}

// prepareRequest insert all pre-defined stuff
func (c *Client) prepareRequest(method, what string, opts map[string]string) (req *http.Request) {
	var endPoint string

	// This is a hack to fetch direct urls for results
	if method == "FETCH" {
		endPoint = what
		method = "GET"
	} else {
		if c.config.endpoint != "" {
			endPoint = fmt.Sprintf("%s/%s", c.config.endpoint, what)
		} else {
			endPoint = fmt.Sprintf("%s/%s", apiEndpoint, what)
		}
	}

	c.mergeGlobalOptions(opts)
	c.verbose("Options:\n%v", opts)
	baseURL := AddQueryParameters(endPoint, opts)

	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		c.log.Printf("error parsing %s: %v", baseURL, err)
		return &http.Request{}
	}

	// We need these when we POST
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
	}

	return
}

// client.handleAPIResponsese check status code & errors from the API
func (c *Client) handleAPIResponsese(r *http.Response) (err error) {
	if r == nil {
		return fmt.Errorf("error: r is nil")
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
			c.log.Printf("Error handling error: %s - %v", r.Body, err)
		}

		c.log.Printf("Info 3XX status: %d code: %d - r:%v\n",
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
		c.log.Printf("Error handling error: %s - %v", r.Body, err)
	}

	err = fmt.Errorf("status: %d code: %d - r:%s\nerrors: %v",
		aerr.Error.Status,
		aerr.Error.Code,
		aerr.Error.Detail,
		aerr.Error.Errors)
	return
}
