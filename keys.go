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
func fetchOneKeyPage(opts map[string]string) (raw *keyList, err error) {

	req := prepareRequest("GET", "keys", opts)

	r, err := ctx.client.Do(req)
	err = handleAPIResponse(r)
	if err != nil {
		return
	}

	raw = &keyList{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, raw)
	//log.Printf("Count=%d raw=%v", raw.Count, r)
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// GetKey returns a given API key
func GetKey(uuid string) (k Key, err error) {
	opts := make(map[string]string)

	req := prepareRequest("GET", fmt.Sprintf("keys/%s", uuid), opts)

	//log.Printf("req: %#v", req)
	r, err := ctx.client.Do(req)
	err = handleAPIResponse(r)
	if err != nil {
		return
	}

	k = Key{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err = json.Unmarshal(body, k)
	//log.Printf("json: %#v\n", p)
	return
}

// GetKeys returns all your API keys
func GetKeys(opts map[string]string) (kl []Key, err error) {

	// First call
	rawlist, err := fetchOneKeyPage(opts)

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

			rawlist, err = fetchOneKeyPage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	kl = res
	return
}
