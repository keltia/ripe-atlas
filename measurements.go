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

// DNS creates a measurement
func DNS(d Definition) (m *Measurement, err error) {
	if checkType(d) || d.Type != "dns" {
		err = ErrInvalidMeasurementType
		return
	}
	return
}

// NTP creates a measurement
func NTP(d Definition) (m *Measurement, err error) {
	return
}

// Ping creates a measurement
func Ping(d Definition) (m *Measurement, err error) {
	return
}

// SSLCert creates a measurement
func SSLCert(d Definition) (m *Measurement, err error) {
	return
}

// Traceroute creates a measurement
func Traceroute(d Definition) (m *Measurement, err error) {
	return
}
