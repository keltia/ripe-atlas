// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"log"
)

// GetProbe returns data for a single probe
func (client *Client) GetProbe(id int) (p *Probe, err error) {
	req := prepareRequest(fmt.Sprintf("probes/%d", id))
	req.Method = rest.Get

	//log.Printf("req: %#v", req)
	r, err := rest.API(req)
	//log.Printf("r: %#v - err: %#v", r, err)
	err = handleAPIResponse(r)
	if err != nil {
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
func fetchOneProbePage(opts map[string]string) (raw *probeList, err error) {

	req := prepareRequest("probes")
	req.Method = rest.Get

	// Do not forget to copy our options
	for qp, val := range opts {
		req.QueryParams[qp] = val
	}

	r, err := rest.API(req)
	err = handleAPIResponse(r)
	if err != nil {
		return
	}

	raw = &probeList{}
	err = json.Unmarshal([]byte(r.Body), raw)
	log.Printf("Count=%d raw=%v", raw.Count, r)
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", raw, r, raw.Next)
	if ctx.config.Verbose {
		fmt.Print("P")
	}
	return
}

// GetProbes returns data for a collection of probes
func (client *Client) GetProbes(opts map[string]string) (p []Probe, err error) {
	// First call
	rawlist, err := fetchOneProbePage(opts)

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

			rawlist, err = fetchOneProbePage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	p = res
	return
}
