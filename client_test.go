package atlas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, c)

	require.IsType(t, (*Client)(nil), c)
	assert.NotEmpty(t, c)
	assert.EqualValues(t, apiEndpoint, c.config.endpoint)
}

func TestNewClient2(t *testing.T) {
	fh := log.New(os.Stderr, "newclient2", log.LstdFlags|log.LUTC)

	c, err := NewClient(Config{Log: fh})
	require.NoError(t, err)
	require.NotNil(t, c)

	require.IsType(t, (*Client)(nil), c)
	assert.NotEmpty(t, c)
	assert.EqualValues(t, apiEndpoint, c.config.endpoint)
	assert.EqualValues(t, fh, c.log)
}

func TestNewClient3(t *testing.T) {
	c, err := NewClient(Config{Verbose: true})
	require.NoError(t, err)
	require.NotNil(t, c)

	require.IsType(t, (*Client)(nil), c)
	assert.NotEmpty(t, c)
	assert.EqualValues(t, apiEndpoint, c.config.endpoint)
	assert.EqualValues(t, 1, c.level)
}

func TestNewClient4(t *testing.T) {
	c, err := NewClient(Config{Level: 2})
	require.NoError(t, err)
	require.NotNil(t, c)

	require.IsType(t, (*Client)(nil), c)
	assert.NotEmpty(t, c)
	assert.EqualValues(t, apiEndpoint, c.config.endpoint)
	assert.EqualValues(t, 2, c.level)
}

func TestNewClient5(t *testing.T) {
	c, err := NewClient(Config{Level: 255})
	require.NoError(t, err)
	require.NotNil(t, c)

	require.IsType(t, (*Client)(nil), c)
	assert.NotEmpty(t, c)
	assert.EqualValues(t, apiEndpoint, c.config.endpoint)
	assert.EqualValues(t, 2, c.level)
}

func TestGetVersion(t *testing.T) {
	ver := GetVersion()
	assert.EqualValues(t, ourVersion, ver, "should be equal")
}

func TestClient_HasAPIKey(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, c)

	key, yes := c.HasAPIKey()
	assert.False(t, yes)
	assert.Empty(t, key)
}

func TestClient_HasAPIKey2(t *testing.T) {
	c, err := NewClient(Config{APIKey: "foo"})
	require.NoError(t, err)
	require.NotNil(t, c)

	key, yes := c.HasAPIKey()
	assert.True(t, yes)
	assert.NotEmpty(t, key)
	assert.EqualValues(t, "foo", key)
}

func TestClient_SetOption(t *testing.T) {
	c, err := NewClient(Config{APIKey: "foobar"})
	require.NoError(t, err)
	require.NotNil(t, c)

	d := c.SetOption("foo", "bar")

	assert.Equal(t, c, d)
	assert.IsType(t, (*Client)(nil), d)

	assert.NotEmpty(t, c.opts)

	_, ok := c.opts["foo"]
	assert.True(t, ok)
}

func TestClient_SetOption2(t *testing.T) {
	c, err := NewClient(Config{APIKey: "foobar"})
	require.NoError(t, err)
	require.NotNil(t, c)

	d := c.SetOption("foo", "")

	assert.Equal(t, c, d)
	assert.IsType(t, (*Client)(nil), d)

	assert.Empty(t, c.opts)

	_, ok := c.opts["foo"]
	assert.False(t, ok)
}

func TestClient_Call(t *testing.T) {
	defer gock.Off()

	d := []Definition{{Type: "foo"}}
	r := &MeasurementRequest{Definitions: d}
	jr, _ := json.Marshal(r)
	//myrp := MeasurementResp{}

	t.Logf("jr=%v", string(jr))

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
		JSON(r).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)
	require.NotNil(t, c)
	require.NotNil(t, c.client)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	opts := map[string]string{"key": "foobar"}
	req := c.prepareRequest("POST", "/measurements/dns/", opts)
	require.NotNil(t, req)

	buf := bytes.NewReader(jr)
	req.Body = ioutil.NopCloser(buf)
	req.ContentLength = int64(buf.Len())

	resp, err := c.call(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 403, resp.StatusCode)
}
