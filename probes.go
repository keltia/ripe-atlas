// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"fmt"
	"github.com/bndr/gopencils"
	"log"
)

// GetProbe returns data for a single probe
func GetProbe(id int) (p *Probe, err error) {
	key, ok := HasAPIKey()
	api := gopencils.Api(apiEndpoint, nil)

	// Add at least one option, the APIkey if present
	var opts = make(map[string]string)

	if ok {
		opts["key"] = key
	}

	r, err := api.Res("probes").Id(id, &p).Get(opts)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}
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
func fetchOneProbePage(api *gopencils.Resource, opts map[string]string) (raw *probeList, err error) {
	r, err := api.Res("probes", &raw).Get(opts)
	if err != nil {
		log.Printf("err: %v", err)
		err = fmt.Errorf("%v - r:%v\n", err, r)
	}
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// GetProbes returns data for a collection of probes
func GetProbes(opts map[string]string) (p []Probe, err error) {
	key, ok := HasAPIKey()
	api := gopencils.Api(apiEndpoint, nil)

	// Add APIKey if set
	if ok {
		opts["key"] = key
	}

	// First call
	rawlist, err := fetchOneProbePage(api, opts)

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

			rawlist, err = fetchOneProbePage(api, opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	p = res
	return
}
