// credits.go
//
// This file implements the credits API calls

package atlas

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// GetCredits returns high-level data for credits
func (c *Client) GetCredits() (credits *Credits, err error) {

	opts := make(map[string]string)
	opts = c.addAPIKey(opts)
	opts = mergeOptions(opts, c.opts)

	req := c.prepareRequest("GET", "credits", opts)

	resp, err := c.call(req)
	c.debug("resp: %#v - err: %#v", resp, err)
	if err != nil {
		return &Credits{}, errors.Wrap(err, "GetCredits/call")
	}

	// We may have all http errors here but the request did succeed
	c.debug("http.code=%d", resp.StatusCode)

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return &Credits{}, errors.Wrapf(err, "GetCredits")
	}

	credits = &Credits{}

	err = json.Unmarshal(body, credits)
	c.debug("json: %#v\n", credits)
	return
}
