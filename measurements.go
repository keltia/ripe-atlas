package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	allTypes = map[string]bool{
		"dns":        true,
		"http":       true,
		"ntp":        true,
		"ping":       true,
		"sslcert":    true,
		"traceroute": true,
		"wifi":       true,
	}
)

// -- private

// checkType verify that the type is valid
func checkType(d Definition) (valid bool) {
	_, ok := allTypes[d.Type]
	return ok
}

// checkTypeAs is a shortcut
func checkTypeAs(d Definition, t string) bool {
	valid := checkType(d)
	return valid && d.Type == t
}

// checkAllTypesAs is a generalization of checkTypeAs
func checkAllTypesAs(dl []Definition, t string) (valid bool) {
	valid = true
	for _, d := range dl {
		if d.Type != t {
			valid = false
			break
		}
	}
	return
}

// measurementList is our main answer
type measurementList struct {
	Count    int
	Next     string
	Previous string
	Results  []Measurement
}

// fetch the given resource
func (c *Client) fetchOneMeasurementPage(opts map[string]string) (raw *measurementList, err error) {
	opts = c.addAPIKey(opts)
	c.mergeGlobalOptions(opts)
	req := c.prepareRequest("GET", "measurements", opts)

	//log.Printf("req=%s qp=%#v", MeasurementEP, opts)
	resp, err := c.call(req)
	if err != nil {
		err = c.handleAPIResponsese(resp)
		if err != nil {
			return
		}
	}
	raw = &measurementList{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, raw)
	//log.Printf("Count=%d raw=%v", raw.Count, resp)
	//log.Printf(">> rawlist=%+v resp=%+v Next=|%s|", rawlist, resp, rawlist.Next)
	return
}

// -- public

// GetMeasurement gets info for a single one
func (c *Client) GetMeasurement(id int) (m *Measurement, err error) {
	opts := make(map[string]string)
	opts = c.addAPIKey(opts)

	c.mergeGlobalOptions(opts)
	req := c.prepareRequest("GET", fmt.Sprintf("measurements/%d", id), opts)

	//log.Printf("req: %#v", req)
	resp, err := c.call(req)
	if err != nil {
		err = c.handleAPIResponsese(resp)
		if err != nil {
			return
		}
	}

	m = &Measurement{}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(body, m)
	//log.Printf("json: %#v\n", m)
	return
}

// DeleteMeasurement stops (not really deletes) a given measurement
func (c *Client) DeleteMeasurement(id int) (err error) {
	opts := make(map[string]string)
	opts = c.addAPIKey(opts)

	req := c.prepareRequest("DELETE", fmt.Sprintf("measurements/%d", id), opts)

	//log.Printf("req: %#v", req)
	resp, err := c.call(req)
	err = c.handleAPIResponsese(resp)
	return
}

// GetMeasurements gets info for a set
func (c *Client) GetMeasurements(opts map[string]string) (m []Measurement, err error) {
	// First call
	rawlist, err := c.fetchOneMeasurementPage(opts)

	// Empty answer
	if rawlist.Count == 0 {
		return []Measurement{}, nil
	}

	var res []Measurement

	res = append(res, rawlist.Results...)
	if rawlist.Next != "" {
		// We have pagination
		for pn := getPageNum(rawlist.Next); rawlist.Next != ""; pn = getPageNum(rawlist.Next) {
			opts["page"] = pn

			rawlist, err = c.fetchOneMeasurementPage(opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	m = res
	return
}
