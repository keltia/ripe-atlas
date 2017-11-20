// credits.go
//
// This file implements the credits API calls

package atlas

import (
	"encoding/json"
	"io/ioutil"
)

// GetCredits returns high-level data for credits
func (client *Client) GetCredits() (credits *Credits, err error) {

	opts := make(map[string]string)
	client.mergeGlobalOptions(opts)

	req := client.prepareRequest("GET", "credits", opts)

	resp, err := client.call(req)
	//log.Printf("resp: %#v - err: %#v", resp, err)
	if err != nil {
		if client.config.Verbose {
			client.log.Printf("API error: %v", err)
		}
		err = handleAPIResponse(resp)
		if err != nil {
			client.log.Printf("error getting credits: %#v", err)
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
