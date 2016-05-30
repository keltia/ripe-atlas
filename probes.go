// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"fmt"
	"github.com/bndr/gopencils"
	"log"
	"regexp"
)

// GetProbe returns data for a single probe
func GetProbe(id int) (p *Probe, err error) {
	auth := WantAuth()
	api := gopencils.Api(apiEndpoint, auth)
	r, err := api.Res("probes").Id(id, &p).Get()
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}
	return
}

// ProbesList is our main answer
type ProbesList struct {
	Count    int
	Next     string
	Previous string
	Results  []Probe
}

// getPageNum returns the value of the page= parameter
func getPageNum(url string) (page string) {
	re := regexp.MustCompile(`page=(\d+)`)
	if m := re.FindStringSubmatch(url); len(m) >= 1 {
		return m[1]
	}
	return ""
}

// GetProbes returns data for a collection of probes
func GetProbes(opts map[string]string) (p []Probe, err error) {
	log.Printf("GetProbes: opts=%+v", opts)
	auth := WantAuth()
	api := gopencils.Api(apiEndpoint, auth)

	var rawlist ProbesList

	r, err := api.Res("probes", &rawlist).Get(opts)
	log.Printf("rawlist=%+v r=%+v", rawlist, r)
	if err != nil {
		err = fmt.Errorf("%v - r:%v\n", err, r)
		return
	}

	// Empty answer
	if rawlist.Count == 0 {
		return nil, fmt.Errorf("empty probe list")
	}

	var res []Probe

	res = append(res, rawlist.Results...)
	if rawlist.Next != "" {
		// We have pagination
		for pn := 2; pn != 0; getPageNum(rawlist.Next) {
			opts["page"] = string(pn)
			r, err = api.Res("probes", &rawlist).Get(opts)
			res = append(res, rawlist.Results...)
		}
	}
	p = res
	fmt.Printf("r: %#v\np: %#v\n", r, p)
	return
}
