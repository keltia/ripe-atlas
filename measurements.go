package atlas

import (
	"errors"
	"github.com/bndr/gopencils"
	"log"
	"fmt"
)

var (
	allTypes = []string{
		"dns",
		"http",
		"ntp",
		"ping",
		"sslcert",
		"traceroute",
		"wifi",
	}
)

// ErrInvalidMeasurementType is a new error
var ErrInvalidMeasurementType = errors.New("invalid measurement type")

var ErrInvalidAPIKey = errors.New("invalid API key")

// -- private

// checkType verify that the type is valid
func checkType(d Definition) (valid bool) {
	valid = false
	for _, t := range allTypes {
		if d.Type == t {
			valid = true
			break
		}
	}
	return
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
func fetchOneMeasurementPage(api *gopencils.Resource, opts map[string]string) (raw *measurementList, err error) {
	r, err := api.Res("measurements", &raw).Get(opts)
	if err != nil {
		log.Printf("err: %v", err)
		err = fmt.Errorf("%v - r:%v\n", err, r)
	}
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// -- public

// GetMeasurement gets info for a single one
func GetMeasurement(id int) (m *Measurement, err error) {
	key, ok := HasAPIKey()
	api := gopencils.Api(apiEndpoint, nil)

	// Add at least one option, the APIkey if present
	var opts map[string]string

	if ok {
		opts["key"] = key
	}

	r, err := api.Res("measurements").Id(id, &m).Get(opts)
	if err != nil {
		err = fmt.Errorf("%v - r:%#v\n", err, r.Raw)
		return
	}
	return
}

// GetMeasurements gets info for a set
func GetMeasurements(opts map[string]string) (m []Measurement, err error) {
	key, ok := HasAPIKey()
	api := gopencils.Api(apiEndpoint, nil)

	// Add at least one option, the APIkey if present
	if ok {
		opts["key"] = key
	}

	rawlist, err := fetchOneMeasurementPage(api, opts)

	// Empty answer
	if rawlist.Count == 0 {
		return nil, fmt.Errorf("empty measurement list")
	}

	var res []Measurement

	res = append(res, rawlist.Results...)
	if rawlist.Next != "" {
		// We have pagination
		for pn := getPageNum(rawlist.Next); rawlist.Next != ""; pn = getPageNum(rawlist.Next) {
			opts["page"] = pn

			rawlist, err = fetchOneMeasurementPage(api, opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	m = res
	return
}

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

// Ping creates a measurement
func Ping(d MeasurementRequest) (m *Measurement, err error) {
	// Check that all Definition.Type are the same and compliant
	if !checkAllTypesAs(d.Definitions, "ping") {
		err = ErrInvalidMeasurementType
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

// Measurement-related methods

// Start is for starting a given measurement
func (m *Measurement) Start(id int) (err error) {
	return nil
}

// Stop is for stopping a given measurement
func (m *Measurement) Stop(id int) (err error) {
	return nil
}
