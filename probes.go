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
	api := gopencils.Api(apiEndpoint, nil)

	p = &Probe{}
	r, err := api.Res("probes", &p).Id(id).Get()
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}
//	fmt.Printf("r: %#v\np: %#v\n", r, p)
	return
}

type Probes []Probe

// GetProbes returns data for a collection of probes
func GetProbes() (p *interface{}, err error) {
	api := gopencils.Api(apiEndpoint, nil)

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
