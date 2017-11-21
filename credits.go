// credits.go
//
// This file implements the credits API calls

package atlas

import (
	"encoding/json"
	"io/ioutil"
)

// GetCredits returns high-level data for credits
func (c *Client) GetCredits() (credits *Credits, err error) {

	opts := make(map[string]string)
	c.mergeGlobalOptions(opts)

	req := c.prepareRequest("GET", "credits", opts)

	resp, err := c.call(req)
	//log.Printf("resp: %#v - err: %#v", resp, err)
	if err != nil {
		if c.config.Verbose {
			c.log.Printf("API error: %v", err)
		}
		err = c.handleAPIResponsese(resp)
		if err != nil {
			c.log.Printf("error getting credits: %#v", err)
			return
		}
	}

	credits = &Credits{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, credits)
	//log.Printf("json: %#v\n", credits)
	return
}
