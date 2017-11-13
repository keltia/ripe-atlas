// keys.go

// This file implements the keys API calls

package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type keyList struct {
	Count    int
	Next     string
	Previous string
	Results  []Key
}

// fetch the given resource
func (client *Client) fetchOneKeyPage(opts map[string]string) (raw *keyList, err error) {

	req := client.prepareRequest("GET", "keys", opts)

	resp, err := client.call(req)
	err = handleAPIResponse(resp)
	if err != nil {
		return
	}

	raw = &keyList{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, raw)
	//log.Printf("Count=%d raw=%v", raw.Count, resp)
	//log.Printf(">> rawlist=%+v resp=%+v Next=|%s|", rawlist, resp, rawlist.Next)
	return
}

// GetKey returns a given API key
func (client *Client) GetKey(uuid string) (k Key, err error) {
	opts := make(map[string]string)

	req := client.prepareRequest("GET", fmt.Sprintf("keys/%s", uuid), opts)

	//log.Printf("req: %#v", req)
	resp, err := client.call(req)
	err = handleAPIResponse(resp)
	if err != nil {
		return
	}

	k = Key{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, k)
	//log.Printf("json: %#v\n", p)
	return
}

// GetKeys returns all your API keys
func (client *Client) GetKeys(opts map[string]string) (kl []Key, err error) {

	// First call
	rawlist, err := client.fetchOneKeyPage(opts)

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

			rawlist, err = client.fetchOneKeyPage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	kl = res
	return
}
