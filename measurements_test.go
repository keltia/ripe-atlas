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

func TestCheckTypeAs(t *testing.T) {
	d := Definition{Type: "dns"}
	test := checkTypeAs(d, "foo")
	if test == true {
		t.Errorf("test should be false")
	}

	test = checkTypeAs(d, "dns")
	if test != true {
		t.Errorf("test should be true: %s", d.Type)
	}
}

func TestCheckAllTypesAs(t *testing.T) {
	dl := []Definition{
		{Type: "foo"},
		{Type: "ping"},
	}

	valid := checkAllTypesAs(dl, "ping")
	if valid != false {
		t.Errorf("valid should be false")
	}

	dl = []Definition{
			{Type: "dns"},
			{Type: "ping"},
		}
	valid = checkAllTypesAs(dl, "ping")
	if valid != false {
		t.Errorf("valid should be false")
	}

	dl = []Definition{
			{Type: "ping"},
			{Type: "ping"},
		}
	valid = checkAllTypesAs(dl, "ping")
	if valid != true {
		t.Errorf("valid should be true")
	}
}

func TestDNS(t *testing.T) {
	d := Definition{Type: "foo"}

	_, err := DNS(d)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}

func TestNTP(t *testing.T) {
	d := Definition{Type: "foo"}

	_, err := NTP(d)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}

func TestPing(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := MeasurementRequest{Definitions:d}

	_, err := Ping(r)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}

func TestSSLCert(t *testing.T) {
	d := Definition{Type: "foo"}

	_, err := SSLCert(d)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}

func TestTraceroute(t *testing.T) {
	d := Definition{Type: "foo"}

	_, err := Traceroute(d)
	if err != ErrInvalidMeasurementType {
		t.Errorf("error %v should be %v", err, ErrInvalidMeasurementType)
	}
}
