package atlas

import (
	"github.com/stretchr/testify/assert"
	"github.com/jarcoal/httpmock"
	"testing"
	"net/http"
	"encoding/json"
	"github.com/sendgrid/rest"
	"bytes"
	"log"
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
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to add a new measurement
	httpmock.RegisterResponder("POST", apiEndpoint+"/measurements/dns/",
		func(req *http.Request) (*http.Response, error) {
			var reqData MeasurementRequest
			var ap APIError
			var body bytes.Buffer
			var badType rest.Response
			var myerr error

			//respData := new(MeasurementResp)

			if err := json.NewDecoder(req.Body).Decode(&reqData); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}

			log.Printf("test.req=%#v", reqData)
			if reqData.Definitions[0].Type != "dns" {

				ap.Error.Status = 500
				ap.Error.Code = 501
				ap.Error.Title = "Bad Type"
				ap.Error.Detail = "Type is not dns"
				badType.StatusCode = 500
				myerr = ErrInvalidMeasurementType
				if err := json.NewEncoder(&body).Encode(ap); err != nil {
					resp, _ := httpmock.NewJsonResponse(500, "argh")
					return resp, myerr
				}
			} else {
				ap.Error.Status = 200
				ap.Error.Code = 200
				ap.Error.Title = "Good Type"
				ap.Error.Detail = "Type is dns"
				badType.StatusCode = 200
				myerr = nil
			}

			if err := json.NewEncoder(&body).Encode(ap); err != nil {
				resp, _ := httpmock.NewJsonResponse(500, "argh")
				return resp, myerr
			}

			resp, err := httpmock.NewJsonResponse(200, body.String())
			if err != nil {
				return httpmock.NewStringResponse(500, ""), myerr
			}
			return resp, myerr
		},
	)


	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	myrp := MeasurementResp{}

	rp, err := DNS(r)
	assert.Error(t, err, "should be in error")
	assert.EqualValues(t, myrp, rp, "should be equal")
	assert.EqualValues(t, ErrInvalidMeasurementType, err, "should be equal")
}
/*
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
*/