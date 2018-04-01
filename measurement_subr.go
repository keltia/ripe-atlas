package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

// MeasurementResp contains all the results of the measurements
type MeasurementResp struct {
	Measurements []int
}

// NewMeasurement create a new MeasurementRequest and fills some fields
func (c *Client) NewMeasurement() (req *MeasurementRequest) {
	var defs []Definition

	req = &MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
		Probes:      *NewProbeSet(c.config.PoolSize, c.config.AreaType, c.config.AreaValue),
	}
	c.verbose("probes: %#v", req.Probes)
	return
}

// NewProbeSet create a set of probes for later requests
func NewProbeSet(howmany int, settype, value string) (ps *ProbeSet) {
	if howmany == 0 {
		howmany = 10
	}

	if settype == "" {
		settype = "area"
	}

	if value == "" {
		value = "WW"
	}

	ps = &ProbeSet{
		{
			Requested: howmany,
			Type:      settype,
			Value:     value,
			Tags:      nil,
		},
	}
	return
}

// SetParams set a few parameters in a definition list
/*
The goal here is to give a dictionary of string and let it figure out each field's type
depending on the recipient's type in the struct.
*/
func (d *Definition) setParams(fields map[string]string) {
	sdef := reflect.ValueOf(d).Elem()
	typeOfDef := sdef.Type()
	for k, v := range fields {
		// Check the field is present
		if f, ok := typeOfDef.FieldByName(k); ok {
			// Use the right type
			switch f.Type.Name() {
			case "float":
				vf, _ := strconv.ParseFloat(v, 32)
				sdef.FieldByName(k).SetFloat(vf)
			case "int":
				vi, _ := strconv.ParseInt(v, 10, 32)
				sdef.FieldByName(k).SetInt(vi)
			case "string":
				sdef.FieldByName(k).SetString(v)
			case "bool":
				vb, _ := strconv.ParseBool(v)
				sdef.FieldByName(k).SetBool(vb)
			default:
				log.Printf("Unsupported type: %s", f.Type.Name())
			}
		}
	}
}

// AddDefinition create a new MeasurementRequest and fills some fields
func (m *MeasurementRequest) AddDefinition(fields map[string]string) *MeasurementRequest {
	def := new(Definition)
	def.setParams(fields)
	m.Definitions = append(m.Definitions, *def)

	return m
}

// createMeasurement creates a measurement for all types
func (c *Client) createMeasurement(t string, d *MeasurementRequest) (m *MeasurementResp, err error) {
	opts := make(map[string]string)
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

	err = c.handleAPIResponsese(resp)
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
