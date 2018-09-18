// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// GetProbe returns data for a single probe
func (c *Client) GetProbe(id int) (p *Probe, err error) {

	opts := make(map[string]string)
	opts = c.addAPIKey(opts)

	c.mergeGlobalOptions(opts)

	req := c.prepareRequest("GET", fmt.Sprintf("probes/%d", id), opts)
	c.debug("req=%#v", req)

	resp, err := c.call(req)
	c.debug("resp: %#v - err: %#v", resp, err)
	if err != nil {
		return &Probe{}, errors.Wrap(err, "GetProbe/call")
	}

	// We may have all http errors here but the request did succeed
	c.debug("http.code=%d", resp.StatusCode)

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return &Probe{}, errors.Wrap(err, "GetProbe")
	}

	p = &Probe{}
	err = json.Unmarshal(body, p)
	c.debug("json: %#v\n", p)
	return
}

// probeList is our main answer
type probeList struct {
	Count    int
	Next     string
	Previous string
	Results  []Probe
}

// fetch the given resource
func (c *Client) fetchOneProbePage(opts map[string]string) (raw *probeList, err error) {

	c.mergeGlobalOptions(opts)
	opts = c.addAPIKey(opts)
	req := c.prepareRequest("GET", "probes", opts)

	resp, err := c.call(req)
	if err != nil {
		if err != nil {
			return &probeList{}, errors.Wrap(err, "fetchOneProbePage/call")
		}
		c.verbose("API error: %v", err)
		_, err = c.handleAPIResponse(resp)
		if err != nil {
			return &probeList{}, errors.Wrap(err, "fetchOneProbePage")
		}
	}

	// We may have all http errors here but the request did succeed
	c.debug("http.code=%d", resp.StatusCode)

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return &probeList{}, errors.Wrap(err, "GetProbes")
	}

	raw = &probeList{}

	err = json.Unmarshal(body, raw)
	if err != nil {
		c.log.Printf("err reading json: raw=%#v err=%v", raw, err)
		return raw, errors.Wrapf(err, "fetchOneProbePage")
	}
	c.verbose("Count=%d raw=%v", raw.Count, resp)
	c.debug("P")
	return
}

// GetProbes returns data for a collection of probes
func (c *Client) GetProbes(opts map[string]string) (p []Probe, err error) {
	// First call
	rawlist, err := c.fetchOneProbePage(opts)
	if err != nil {
		return []Probe{}, errors.Wrap(err, "GetProbes")
	}

	// Empty answer
	if rawlist.Count == 0 {
		return nil, fmt.Errorf("empty probe list")
	}

	var res []Probe

	res = append(res, rawlist.Results...)
	if rawlist.Next != "" {
		// We have pagination
		for pn := getPageNum(rawlist.Next); rawlist.Next != ""; pn = getPageNum(rawlist.Next) {
			opts["page"] = pn

			rawlist, err = c.fetchOneProbePage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	p = res
	return
}
