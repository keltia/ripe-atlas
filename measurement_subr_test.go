package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		Log:          nil,
	}
)

func Before(t *testing.T) *Client {

	testc, err := NewClient(TesCfg)

	assert.NoError(t, err)
	assert.NotNil(t, testc)
	assert.IsType(t, (*Client)(nil), testc)
	require.NotNil(t, testc.client)

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

func TestClient_Call(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myurl, _ := url.Parse(apiEndpoint)

	gock.New(apiEndpoint).
		Post("measurements/dns").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		MatchType("json").
		JSON(r).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	opts := map[string]string{"key": "foobar"}
	req := c.prepareRequest("POST", "/measurements/dns", opts)
	require.NotNil(t, req)

	buf := bytes.NewReader(jr)
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	require.NotNil(t, req.Body)

	resp, err := c.call(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClient_DNS_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

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
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.DNS(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_DNS(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

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
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.DNS(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}

func TestClient_NTP_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

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
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.NTP(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_NTP(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

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
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.NTP(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}

func TestClient_Ping_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/ping").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.Ping(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_Ping(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/ping").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.Ping(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}

func TestClient_Traceroute_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/traceroute").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.Traceroute(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_Traceroute(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/traceroute").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.Traceroute(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}

func TestClient_HTTP_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/http").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.HTTP(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_HTTP(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/http").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.HTTP(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}

func TestClient_SSLCert_InvalidKey(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/sslcert").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(403).
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.SSLCert(r)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.EqualValues(t, "createMeasurement: The provided API key does not exist", err.Error())
}

func TestClient_SSLCert(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	jrq, _ := json.Marshal(MeasurementResp{})

	myurl, _ := url.Parse(apiEndpoint)

	buf := bytes.NewReader(jr)
	gock.New(apiEndpoint).
		Post("measurements/sslcert").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"content-type": "application/json",
			"accept":       "application/json",
			"host":         myurl.Host,
			"user-agent":   fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Body(buf).
		Reply(200).
		BodyString(string(jrq))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.SSLCert(r)
	assert.NoError(t, err)
	assert.Empty(t, rp)
}
