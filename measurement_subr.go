package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"log"
	"reflect"
	"strconv"
)

// MeasurementResp contains all the results of the measurements
type MeasurementResp struct {
	Measurements []int
}

var (
	// If nothing is specified, use this
	defProbeSet = ProbeSet{
		{
			Requested: 10,
			Type:      "area",
			Value:     "WW",
			Tags:      nil,
		},
	}
)

// NewMeasurement create a new MeasurementRequest and fills some fields
func NewMeasurement() (req *MeasurementRequest) {
	var defs []Definition

	req = &MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
		Probes:      defProbeSet,
	}
	return
}

// NewProbeSet create a set of probes for later requests
func NewProbeSet(howmany int) (ps *ProbeSet) {
	ps = &ProbeSet{
		{
			Requested: howmany,
			Type:      "area",
			Value:     "WW",
			Tags:      nil,
		},
	}
	return
}

// SetParams set a few parameters in a definition list
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
func createMeasurement(t string, d *MeasurementRequest) (m *MeasurementResp, err error) {
	req := prepareRequest(fmt.Sprintf("measurements/%s", t))

	body, err := json.Marshal(d)
	if err != nil {
		return
	}

	req.Method = rest.Post
	req.Body = body

	log.Printf("body: %s", body)
	resp, err := rest.API(req)
	log.Printf("resp: %v", resp)
	if err != nil {
		log.Printf("err: %v", err)
		//return
	}

	err = handleAPIResponse(resp)
	if err != nil {
		return
	}

	m = &MeasurementResp{}
	err = json.Unmarshal([]byte(resp.Body), m)
	//r, err := api.Res(base, &resp).Post(d)
	fmt.Printf("m: %v\nresp: %#v\nd: %v\n", m, string(resp.Body), d)
	if err != nil {
		err = fmt.Errorf("err: %v - m:%v", err, m)
		return
	}

	return
}

// DNS creates a measurement
func DNS(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("dns", d)
}

// HTTP creates a measurement
func HTTP(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("http", d)
}

// NTP creates a measurement
func NTP(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("ntp", d)
}

// Ping creates a measurement
func Ping(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("ping", d)
}

// SSLCert creates a measurement
func SSLCert(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("sslcert", d)
}

// Traceroute creates a measurement
func Traceroute(d *MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("traceroute", d)
}
