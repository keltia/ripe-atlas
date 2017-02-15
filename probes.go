// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
)

// GetProbe returns data for a single probe
func GetProbe(id int) (p *Probe, err error) {
	probeEP := apiEndpoint + "/probes"

	key, ok := HasAPIKey()

	// Add at least one option, the APIkey if present
	hdrs := make(map[string]string)
	opts := make(map[string]string)

	if ok {
		opts["key"] = key
	}

	req := rest.Request{
		BaseURL:     probeEP + fmt.Sprintf("/%d", id),
		Method:      rest.Get,
		Headers:     hdrs,
		QueryParams: opts,
	}

	//log.Printf("req: %#v", req)
	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	p = &Probe{}
	err = json.Unmarshal([]byte(r.Body), p)
	//log.Printf("json: %#v\n", p)
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
func fetchOneProbePage(probeEP string, opts map[string]string) (raw *probeList, err error) {
	hdrs := make(map[string]string)
	req := rest.Request{
		BaseURL:     probeEP,
		Method:      rest.Get,
		Headers:     hdrs,
		QueryParams: opts,
	}

	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	raw = &probeList{}
	err = json.Unmarshal([]byte(r.Body), raw)
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// GetProbes returns data for a collection of probes
func GetProbes(opts map[string]string) (p []Probe, err error) {
	probeEP := apiEndpoint + "/probes"

	key, ok := HasAPIKey()

	// Add APIKey if set
	if ok {
		opts["key"] = key
	}

	// First call
	rawlist, err := fetchOneProbePage(probeEP, opts)

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

			rawlist, err = fetchOneProbePage(probeEP, opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	p = res
	return
}
