package atlas

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/goware/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	TesCfg = Config{
		APIKey:       "foobar",
		DefaultProbe: 666,
		IsOneOff:     true,
		PoolSize:     10,
		AreaType:     "country",
		AreaValue:    "fr",
		Tags:         "",
		Verbose:      false,
		Log:          nil,
		endpoint:     testURL,
	}

	mockService *httpmock.MockHTTPServer
)

func Before(t *testing.T) *Client {

	testc, err := NewClient(TesCfg)

	assert.NoError(t, err)
	assert.NotNil(t, testc)
	assert.IsType(t, (*Client)(nil), testc)

	return testc
}

func TestClient_NewMeasurement(t *testing.T) {
	c := Before(t)

	m := c.NewMeasurement()

	assert.NotNil(t, c)
	assert.NotNil(t, m)
	assert.IsType(t, (*MeasurementRequest)(nil), m)
	assert.IsType(t, ([]Definition)(nil), m.Definitions)
	assert.IsType(t, ([]ProbeSet)(nil), m.Probes)
	assert.True(t, m.IsOneoff)
}

func TestMeasurementRequest_AddDefinition(t *testing.T) {
	c := Before(t)
	require.NotNil(t, c)
	require.NotEmpty(t, c)

	mr := c.NewMeasurement()
	require.NotNil(t, mr)
	assert.IsType(t, (*MeasurementRequest)(nil), mr)

	mrlen := len(mr.Definitions)

	opts := map[string]string{"AF": "6"}

	mrr := mr.AddDefinition(opts)
	require.NotNil(t, mrr)
	assert.IsType(t, (*MeasurementRequest)(nil), mrr)

	assert.Equal(t, mrlen+1, len(mrr.Definitions))
}

func TestIsPositive(t *testing.T) {
	a := ""
	b, y := isPositive(a)
	assert.True(t, y)
	assert.Equal(t, "", b)

	a = "foo"
	b, y = isPositive(a)
	assert.True(t, y)
	assert.Equal(t, "foo", b)

	a = "+foo"
	b, y = isPositive(a)
	assert.True(t, y)
	assert.Equal(t, "foo", b)

	a = "-foo"
	b, y = isPositive(a)
	assert.False(t, y)
	assert.Equal(t, "foo", b)

	a = "!foo"
	b, y = isPositive(a)
	assert.False(t, y)
	assert.Equal(t, "foo", b)
}

var TestSplitTagsData = []struct {
	tags    string
	in, out string
}{
	{"", "", ""},
	{"foo", "foo", ""},
	{"foo,bar", "foo,bar", ""},
	{"!foo", "", "foo"},
	{"foo,!bar", "foo", "bar"},
	{"+foo,bar", "foo,bar", ""},
	{"+foo,-bar", "foo", "bar"},
	{"+foo,-bar,!baz", "foo", "bar,baz"},
}

func TestSplitTags(t *testing.T) {
	for _, d := range TestSplitTagsData {
		in, out := splitTags(d.tags)
		assert.Equal(t, d.in, in)
		assert.Equal(t, d.out, out)
	}
}

func TestNewProbeSet(t *testing.T) {
	bmps := ProbeSet{Requested: 10, Type: "country", Value: "fr", TagsInclude: "system-ipv6-stable-1d", TagsExclude: ""}
	ps := NewProbeSet(10, "country", "fr", "system-ipv6-stable-1d")

	assert.NotEmpty(t, ps)
	assert.EqualValues(t, bmps, ps)
}

func TestNewProbeSet_2(t *testing.T) {
	bmps := ProbeSet{Requested: 10, Type: "area", Value: "WW", TagsInclude: "system-ipv6-stable-1d", TagsExclude: ""}
	ps := NewProbeSet(0, "", "", "system-ipv6-stable-1d")

	assert.NotEmpty(t, ps)
	assert.EqualValues(t, bmps, ps)
}

func BeforeAPI(t *testing.T) {
	if mockService == nil {
		// new mocking server
		t.Log("starting mock...")
		mockService = httpmock.NewMockHTTPServer("localhost:10000")
	}

	require.NotNil(t, mockService)

}

func TestDNS(t *testing.T) {
	c := Before(t)
	BeforeAPI(t)

	c.config.Verbose = true

	require.NotNil(t, c)
	require.NotEmpty(t, c)

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	myrp := MeasurementResp{}

	t.Logf("jr=%v", string(jr))

	buf := bytes.NewReader([]byte(jr))

	request1, _ := url.Parse(testURL + "/measurements/dns?key=foobar")
	resp := httpmock.MockResponse{
		Request: http.Request{
			Method: "POST",
			URL:    request1,
			Body:   ioutil.NopCloser(buf),
			ContentLength: int64(buf.Len()),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
				"Accept": []string{"application/json"},
			},
		},
		Response: httpmock.Response{
			StatusCode: 200,
			Body:       string("baz"),
		},
	}
	mockService.AddResponse(resp)
	t.Logf("respmap=%v", mockService.ResponseMap)

	rp, err := c.DNS(r)
	t.Logf("rp=%#v", rp)
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
