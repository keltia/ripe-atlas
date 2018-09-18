package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/h2non/gock"
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
		Verbose:      true,
		Log:          nil,
	}
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

func TestClient_DNS(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	t.Logf("jr=%v", string(jr))

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/dns").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	myerr := "status: 403 code: 104 - r:The provided API key does not exist\nerrors: []"

	rp, err := c.DNS(r)

	t.Logf("rp=%#v", rp)
	assert.Error(t, err)
	assert.Nil(t, rp)
	assert.Equal(t, myerr, err.Error())
}

func TestClient_NTP(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "ntp"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	t.Logf("jr=%v", string(jr))

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/ntp").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	myerr := "status: 403 code: 104 - r:The provided API key does not exist\nerrors: []"

	rp, err := c.NTP(r)

	t.Logf("rp=%#v", rp)
	assert.Error(t, err)
	assert.Nil(t, rp)
	assert.Equal(t, myerr, err.Error())
}

/*
func (c *Client) call(req *http.Request) (*http.Response, error) {
	c.verbose("Full URL:\n%v", req.URL)

	myurl, _ := url.Parse(apiEndpoint)
	req.Header.Set("Host", myurl.Host)
	req.Header.Set("User-Agent", fmt.Sprintf("ripe-atlas/%s", ourVersion))

	return c.client.Do(req)
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
*/
