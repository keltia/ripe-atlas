package atlas

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ResultsResp contains all the results of the measurements
type ResultsResp struct {
	Results []MeasurementResult
}

// GetResults gets results info for a single Measurement ID
func (c *Client) GetResults(id int) (r *ResultsResp, err error) {
	r = &ResultsResp{}

	m, err := c.GetMeasurement(id)
	if err != nil {
		return r, err
	}

	if m.Result == "" {
		return r, nil
	}

	body, err := c.FetchResult(m.Result)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal([]byte(body), &r.Results)

	return
}

// FetchResult downloads result for a given measurement
func (c *Client) FetchResult(url string) (string, error) {
	opts := make(map[string]string)

	c.mergeGlobalOptions(opts)

	// Remove our key for fetching the results
	delete(opts, "key")

	req := c.prepareRequest("FETCH", url, opts)

	c.debug("req=%#v", req)
	c.debug("url=%s", req.URL.String())

	resp, err := c.call(req)
	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return "", errors.Wrap(err, "FetchResult")
	}

	return string(body), nil
}
