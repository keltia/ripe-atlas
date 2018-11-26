// keys.go

// This file implements the keys API calls

package atlas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type keyList struct {
	Count    int
	Next     string
	Previous string
	Results  []Key
}

// fetch the given resource
func (c *Client) fetchOneKeyPage(opts map[string]string) (raw *keyList, err error) {

	opts = c.addAPIKey(opts)
	opts = mergeOptions(opts, c.opts)

	req := c.prepareRequest("GET", "keys", opts)

	resp, err := c.call(req)
	if err != nil {
		return nil, errors.Wrap(err, "fetchOneKeyPage/call")
	}

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return &keyList{}, errors.Wrap(err, "fetchOneKeyPage")
	}

	raw = &keyList{}

	err = json.Unmarshal(body, raw)
	return
}

// GetKey returns a given API key
func (c *Client) GetKey(uuid string) (k Key, err error) {
	opts := make(map[string]string)
	opts = c.addAPIKey(opts)
	opts = mergeOptions(opts, c.opts)

	req := c.prepareRequest("GET", fmt.Sprintf("keys/%s", uuid), opts)

	//log.Printf("req: %#v", req)
	resp, err := c.call(req)
	if err != nil {
		return Key{}, errors.Wrap(err, "GetKey/call")
	}

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return Key{}, errors.Wrap(err, "GetKey")
	}

	k = Key{}

	err = json.Unmarshal(body, &k)
	c.debug("json: %#v\n", k)
	return
}

// GetKeys returns all your API keys
func (c *Client) GetKeys(opts map[string]string) (kl []Key, err error) {

	opts = c.addAPIKey(opts)
	opts = mergeOptions(opts, c.opts)

	// First call
	rawlist, err := c.fetchOneKeyPage(opts)
	if err != nil {
		return []Key{}, errors.Wrap(err, "GetKeys")
	}

	// Empty answer
	if rawlist.Count == 0 {
		return nil, fmt.Errorf("empty key list")
	}

	var res []Key

	res = append(res, rawlist.Results...)
	if rawlist.Next != "" {
		// We have pagination
		for pn := getPageNum(rawlist.Next); rawlist.Next != ""; pn = getPageNum(rawlist.Next) {
			opts["page"] = pn

			rawlist, err = c.fetchOneKeyPage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	kl = res
	return
}
