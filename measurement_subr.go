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

// NewMeasurement create a new MeasurementRequest and fills some fields
func NewMeasurement(t string, fields map[string]string) (req *MeasurementRequest) {
	var defs []Definition

	def := NewDefinition(t, fields)
	probes := ProbeSet{
		{
			Requested: 10,
			Type:      "area",
			Value:     "WW",
			Tags:      nil,
		},
	}
	defs = append(defs, *def)
	req = &MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
		Probes:      probes,
	}
	return
}

// NewDefinition create a new MeasuremenrRequest and fills some fields
func NewDefinition(t string, fields map[string]string) (def *Definition) {
	def = &Definition{
		Type: t,
	}
	sdef := reflect.ValueOf(&def).Elem()
	typeOfDef := sdef.Type()
	for k, v := range fields {
		// Check the field is present
		if f, ok := typeOfDef.FieldByName(k); ok {
			// Use the right type
			switch f.Name {
			case "float":
				vf, _ := strconv.ParseFloat(v, 32)
				sdef.FieldByName(k).SetFloat(vf)
			case "int":
				vi, _ := strconv.ParseInt(v, 10, 32)
				sdef.FieldByName(k).SetInt(vi)
			case "string":
				sdef.FieldByName(k).SetString(v)
			}
		}
	}
	return
}

// createMeasurement creates a measurement for all types
func createMeasurement(t string, d MeasurementRequest) (m *MeasurementResp, err error) {
	req := prepareRequest(fmt.Sprintf("measurements/%s", t))

	body, err := json.Marshal(d)
	if err != nil {
		return
	}

	req.Method = rest.Post
	req.Body = body

	log.Printf("body: %s", body)
	resp, err := rest.API(req)
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
func DNS(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("dns", d)
}

// HTTP creates a measurement
func HTTP(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("http", d)
}

// NTP creates a measurement
func NTP(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("ntp", d)
}

// Ping creates a measurement
func Ping(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("ping", d)
}

// SSLCert creates a measurement
func SSLCert(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("sslcert", d)
}

// Traceroute creates a measurement
func Traceroute(d MeasurementRequest) (m *MeasurementResp, err error) {
	return createMeasurement("traceroute", d)
}
