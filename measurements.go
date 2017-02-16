package atlas

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/rest"
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
func fetchOneMeasurementPage(MeasurementEP string, opts map[string]string) (raw *measurementList, err error) {
	hdrs := make(map[string]string)
	req := rest.Request{
		BaseURL:     MeasurementEP,
		Method:      rest.Get,
		Headers:     hdrs,
		QueryParams: opts,
	}

	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	raw = &measurementList{}
	err = json.Unmarshal([]byte(r.Body), raw)
	//log.Printf(">> rawlist=%+v r=%+v Next=|%s|", rawlist, r, rawlist.Next)
	return
}

// -- public

// GetMeasurement gets info for a single one
func GetMeasurement(id int) (m *Measurement, err error) {
	measurementEP := apiEndpoint + "/measurements"

	key, ok := HasAPIKey()

	// Add at least one option, the APIkey if present
	hdrs := make(map[string]string)
	opts := make(map[string]string)

	if ok {
		opts["key"] = key
	}

	req := rest.Request{
		BaseURL:     measurementEP + fmt.Sprintf("/%d", id),
		Method:      rest.Get,
		Headers:     hdrs,
		QueryParams: opts,
	}

	//log.Printf("req: %#v", req)
	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
		return
	}

	m = &Measurement{}
	err = json.Unmarshal([]byte(r.Body), m)
	//log.Printf("json: %#v\n", m)
	return
}

// DeleteMeasurement stops (not really deletes) a given measurement
func DeleteMeasurement(id int) (err error) {
	measurementEP := apiEndpoint + "/measurements"

	key, ok := HasAPIKey()

	// Add at least one option, the APIkey if present
	hdrs := make(map[string]string)
	opts := make(map[string]string)

	if ok {
		opts["key"] = key
	}

	req := rest.Request{
		BaseURL:     measurementEP + fmt.Sprintf("/%d", id),
		Method:      rest.Delete,
		Headers:     hdrs,
		QueryParams: opts,
	}

	//log.Printf("req: %#v", req)
	r, err := rest.API(req)
	if err != nil {
		err = fmt.Errorf("err: %v - r:%v\n", err, r)
	}
	return
}

// GetMeasurements gets info for a set
func GetMeasurements(opts map[string]string) (m []Measurement, err error) {
	measurementEP := apiEndpoint + "/measurements/"

	key, ok := HasAPIKey()

	// Add APIKey if set
	if ok {
		opts["key"] = key
	}

	// First call
	rawlist, err := fetchOneMeasurementPage(measurementEP, opts)

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

			rawlist, err = fetchOneMeasurementPage(measurementEP, opts)
			if err != nil {
				return
			}

			res = append(res, rawlist.Results...)
		}
	}
	m = res
	return
}

// Measurement-related methods

// Start is for starting a given measurement
func (m *Measurement) Start(id int) (err error) {
	return nil
}
