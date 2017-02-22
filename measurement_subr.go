package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
	"log"
)

// MeasurementResp contains all the results of the measurements
type MeasurementResp struct {
	Measurements []int
}

// createMeasurement creates a measurement for all types
func createMeasurement(t string, d MeasurementRequest) (m *MeasurementResp, err error) {
    // Check that all Definition.Type are the same and compliant
    if !checkAllTypesAs(d.Definitions, t) {
        err = ErrInvalidMeasurementType
        return
    }

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
        err = fmt.Errorf("err: %v - m:%v\n", err, m)
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
