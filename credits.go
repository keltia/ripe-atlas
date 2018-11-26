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
	c.mergeGlobalOptions(opts)

	req := c.prepareRequest("GET", "credits", opts)

	resp, err := c.call(req)
	if err != nil {
		c.verbose("call: %v", err)
		return &Credits{}, errors.Wrap(err, "call")
	}

	body, err := c.handleAPIResponse(resp)
	if err != nil {
		return &Credits{}, errors.Wrap(err, "GetCredits")
	}

	credits = &Credits{}
	err = json.Unmarshal(body, credits)
	c.debug("credits=%#v\n", credits)
	return
}
