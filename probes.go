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

// GetProbes returns data for a collection of probes
func GetProbes(opts map[string]string) (p []Probe, err error) {
	log.Printf("GetProbes: opts=%+v", opts)
	auth := WantAuth()
	api := gopencils.Api(apiEndpoint, auth)

	var rawlist *ProbesList

	r, err := api.Res("probes", rawlist).Get(opts)
	log.Printf("rawlist=%+v r=%+v", rawlist, r)
	if err != nil {
		err = fmt.Errorf("%v - r:%v\n", err, r)
		return
	}

	// Empty answer
	if rawlist == nil {
		return nil, fmt.Errorf("empty probe list")
	}

	if rawlist.Next != "" {
		// We have pagination

	}
	p = rawlist.Results
	fmt.Printf("r: %#v\np: %#v\n", r, p)
	return
}
