package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// MeasurementResp contains all the results of the measurements
type MeasurementResp struct {
	Measurements []int
}

// NewMeasurement create a new MeasurementRequest and fills some fields
func (c *Client) NewMeasurement() (req *MeasurementRequest) {
	var defs []Definition

	ps := NewProbeSet(c.config.PoolSize, c.config.AreaType, c.config.AreaValue, c.config.Tags)
	req = &MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
		Probes:      []ProbeSet{ps},
	}
	c.verbose("probes: %#v", req.Probes)
	return
}

func isPositive(tag string) (string, bool) {
	bare := tag

	if tag == "" {
		return "", true
	}

	if tag[0] == '+' || tag[0] == '-' || tag[0] == '!' {
		bare = tag[1:]
	}

	if tag[0] == '-' || tag[0] == '!' {
		return bare, false
	}
	return bare, true
}

// splitTags analyse tags values:
//   +tag / tag  ==> tags_include
//   -tag / !tag ==> tags_exclude
func splitTags(tags string) (in, out string) {
	var (
		aIn  []string
		aOut []string
	)

	all := strings.Split(tags, ",")
	if len(all) == 0 {
		return "", ""
	}

	for _, tag := range all {
		if bare, yes := isPositive(tag); yes {
			aIn = append(aIn, bare)
		} else {
			aOut = append(aOut, bare)
		}
	}
	return strings.Join(aIn, ","), strings.Join(aOut, ",")
}

// NewProbeSet create a set of probes for later requests
func NewProbeSet(howmany int, settype, value string, tags string) (ps ProbeSet) {
	var aIn, aOut string

	if howmany == 0 {
		howmany = 10
	}

	if settype == "" {
		settype = "area"
	}

	if value == "" {
		value = "WW"
	}

	// If tags were specified, analyze them
	if tags != "" {
		aIn, aOut = splitTags(tags)
	}

	ps = ProbeSet{
		Requested:   howmany,
		Type:        settype,
		Value:       value,
		TagsInclude: aIn,
		TagsExclude: aOut,
	}
	return
}

// AddDefinition create a new MeasurementRequest and fills some fields
func (m *MeasurementRequest) AddDefinition(fields map[string]string) *MeasurementRequest {
	def := new(Definition)
	err := FillDefinition(def, fields)
	if err == nil {
		m.Definitions = append(m.Definitions, *def)
	}
	return m
}

// createMeasurement creates a measurement for all types
func (c *Client) createMeasurement(t string, d *MeasurementRequest) (m *MeasurementResp, err error) {
	opts := make(map[string]string)
	opts = c.addAPIKey(opts)
	req := c.prepareRequest("POST", fmt.Sprintf("measurements/%s", t), opts)

	body, err := json.Marshal(d)
	if err != nil {
		return
	}

	buf := bytes.NewReader(body)
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	c.verbose("req: %#v", req)
	c.verbose("body: %s", body)
	resp, err := c.call(req)
	c.verbose("resp: %v", resp)
	if err != nil {
		c.log.Printf("err: %v", err)
		//return
	}

	err = c.handleAPIResponse(resp)
	if err != nil {
		return
	}

	m = &MeasurementResp{}
	rbody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	err = json.Unmarshal(rbody, m)
	// Only display if debug/verbose
	c.verbose("m: %v\nresp: %#v\nd: %v\n", m, string(rbody), d)
	if err != nil {
		err = fmt.Errorf("err: %v - m:%v", err, m)
		return
	}

	return
}

// DNS creates a measurement
func (c *Client) DNS(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("dns", d)
}

// HTTP creates a measurement
func (c *Client) HTTP(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("http", d)
}

// NTP creates a measurement
func (c *Client) NTP(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("ntp", d)
}

// Ping creates a measurement
func (c *Client) Ping(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("ping", d)
}

// SSLCert creates a measurement
func (c *Client) SSLCert(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("sslcert", d)
}

// Traceroute creates a measurement
func (c *Client) Traceroute(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return c.createMeasurement("traceroute", d)
}
