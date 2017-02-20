// keys.go

// This file implements the keys API calls

package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
)

type keyList struct {
	Count    int
	Next     string
	Previous string
	Results  []Key
}

// fetch the given resource
func fetchOneKeyPage(opts map[string]string) (raw *keyList, err error) {

	req := prepareRequest("keys")
	req.Method = rest.Get

	// Do not forget to copy our options
	for qp, val := range opts {
		req.QueryParams[qp] = val
	}

	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	raw = &keyList{}
	err = json.Unmarshal([]byte(r.Body), raw)
	//log.Printf("Count=%d raw=%v", raw.Count, r)
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// GetKey returns a given API key
func GetKey(uuid string) (k Key, err error) {

	req := prepareRequest(fmt.Sprintf("keys/%d", uuid))
	req.Method = rest.Get

	//log.Printf("req: %#v", req)
	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	k = Key{}
	err = json.Unmarshal([]byte(r.Body), &k)
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
