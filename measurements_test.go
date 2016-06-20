package atlas

import (
	"testing"
)

func TestCheckType(t *testing.T) {
	d := Definition{Type: "foo"}

	test := checkType(d)
	if test != false {
		t.Errorf("type is invalid: %s", d.Type)
	}

	d = Definition{Type: "dns"}
	test = checkType(d)
	if test != true {
		t.Errorf("type is invalid: %s", d.Type)
	}
}

func TestDNS(t *testing.T) {
	d := Definition{Type: "foo"}

	_, err := DNS(d)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}
