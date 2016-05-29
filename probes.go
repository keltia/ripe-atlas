// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"fmt"
	"github.com/bndr/gopencils"
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

type Probes []Probe

// GetProbes returns data for a collection of probes
func GetProbes(opts map[string]string) (p *ProbesList, err error) {
	log.Printf("GetProbes: opts=%+v", opts)
	auth := WantAuth()
	api := gopencils.Api(apiEndpoint, auth)

	var plist *interface{}

	r, err := api.Res("probes", plist).Get()
	if err != nil {
		err = fmt.Errorf("%v - r:%v\n", err, r)
		return
	}
	fmt.Printf("r: %#v\nplist: %#v\n", r, plist)
	p = plist
	return
}
