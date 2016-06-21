package atlas

import (
	"errors"
)

var (
	allTypes = []string{
		"dns",
		"ntp",
		"ping",
		"sslcert",
		"traceroute",
		"wifi",
	}
)

var ErrInvalidMeasurementType = errors.New("invalid measurement type")

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
func checkTypeAs(d Definition, t string) (valid bool) {
	valid = true
	if checkType(d) && d.Type != t {
		valid = false
	}
	return
}

// DNS creates a measurement
func DNS(d Definition) (m *Measurement, err error) {
	if checkTypeAs(d, "dns") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// NTP creates a measurement
func NTP(d Definition) (m *Measurement, err error) {
	if checkTypeAs(d, "ntp") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// Ping creates a measurement
func Ping(d Definition) (m *Measurement, err error) {
	if checkTypeAs(d, "ping") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// SSLCert creates a measurement
func SSLCert(d Definition) (m *Measurement, err error) {
	if checkTypeAs(d, "sslcert") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// Traceroute creates a measurement
func Traceroute(d Definition) (m *Measurement, err error) {
	if checkTypeAs(d, "traceroute") {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// Measurement-related methods

func (m *Measurement) Start() (err error) {
	return nil
}

func (m *Measurement) Stop() (err error) {
	return nil
}
