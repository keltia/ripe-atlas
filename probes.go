// probes.go
//
// This file implements the probe API calls

package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GetProbe returns data for a single probe
func (client *Client) GetProbe(id int) (p *Probe, err error) {

	opts := make(map[string]string)
	client.mergeGlobalOptions(opts)

	req := client.prepareRequest("GET", fmt.Sprintf("probes/%d", id), opts)

	resp, err := client.call(req)
	//log.Printf("resp: %#v - err: %#v", resp, err)
	if err != nil {
		if client.config.Verbose {
			client.log.Printf("API error: %v", err)
		}
		err = handleAPIResponse(resp)
		if err != nil {
			return
		}
	}

	p = &Probe{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, p)
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
func (client *Client) fetchOneProbePage(opts map[string]string) (raw *probeList, err error) {

	client.mergeGlobalOptions(opts)
	req := client.prepareRequest("GET", "probes", opts)

	resp, err := client.call(req)
	if err != nil {
		if client.config.Verbose {
			client.log.Printf("API error: %v", err)
		}
		err = handleAPIResponse(resp)
		if err != nil {
			return
		}
	}

	raw = &probeList{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, raw)
	if err != nil {
		client.log.Printf("err reading json: raw=%#v err=%v", raw, err)
		return
	}
	if client.config.Verbose {
		client.log.Printf("Count=%d raw=%v", raw.Count, resp)
		fmt.Print("P")
	}
	return
}

// GetProbes returns data for a collection of probes
func (client *Client) GetProbes(opts map[string]string) (p []Probe, err error) {
	// First call
	rawlist, err := client.fetchOneProbePage(opts)

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

			rawlist, err = client.fetchOneProbePage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	p = res
	return
}
