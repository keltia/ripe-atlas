package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckType(t *testing.T) {
	d := Definition{Type: "foo"}

	valid := checkType(d)
	assert.EqualValues(t, false, valid, "should be false")

	d = Definition{Type: "dns"}
	valid = checkType(d)
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckTypeAs(t *testing.T) {
	d := Definition{Type: "dns"}
	valid := checkTypeAs(d, "foo")
	assert.EqualValues(t, false, valid, "should be false")

	valid = checkTypeAs(d, "dns")
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckAllTypesAs(t *testing.T) {
	dl := []Definition{
		{Type: "foo"},
		{Type: "ping"},
	}

	valid := checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "dns"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "ping"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, true, valid, "should be true")
}

func TestDNS(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}

	_, err := DNS(r)
	assert.Error(t, err, "should be an error")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}

func TestNTP(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}

	_, err := NTP(r)
	assert.Error(t, err, "should be an error")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}

func TestPing(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}

	_, err := Ping(r)
	assert.Error(t, err, "should be an error")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}

func TestSSLCert(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}

	_, err := SSLCert(r)
	assert.Error(t, err, "should be an error")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}

func TestTraceroute(t *testing.T) {
	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}

	_, err := Traceroute(r)
	assert.Error(t, err, "should be an error")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}
