package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
)

// DNS creates a measurement
func DNS(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "dns") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// HTTP creates a measurement
func HTTP(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "http") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// NTP creates a measurement
func NTP(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "ntp") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

type PingResp struct {
	Measurements []int
}

type pingError struct {
	Error struct {
		Status int
		Code   int
		Detail string
		Title  string
	}
}

// Ping creates a measurement
func Ping(d MeasurementRequest) (m *PingResp, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "ping") {
		err = ErrInvalidMeasurementType
		return
	}

	pingEP := apiEndpoint + "/measurements/ping"

	// Add at least one option, the APIkey if present
	hdrs := make(map[string]string)
	opts := make(map[string]string)

	key, ok := HasAPIKey()
	if ok {
		opts["key"] = key
	} else {
		err = ErrInvalidAPIKey
		return
	}

	body, err := json.Marshal(d)
	if err != nil {
		return
	}

	req := rest.Request{
		BaseURL:     pingEP,
		Method:      rest.Post,
		Headers:     hdrs,
		QueryParams: opts,
		Body:        body,
	}

	resp, err := rest.API(req)

	m = &PingResp{}
	err = json.Unmarshal([]byte(resp.Body), m)
	//r, err := api.Res(base, &resp).Post(d)
	fmt.Printf("m: %v\nresp: %#v\nd: %v\n", m, string(resp.Body), d)
	if err != nil {
		err = fmt.Errorf("err: %v - m:%v\n", err, m)
		return
	}
	return
}

// SSLCert creates a measurement
func SSLCert(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "sslcert") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// Traceroute creates a measurement
func Traceroute(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "traceroute") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}
