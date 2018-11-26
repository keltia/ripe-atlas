package atlas

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testURL = "http://localhost:10000"

func TestGetPageNum(t *testing.T) {
	nurl := "https://foo.example.com/"

	n := getPageNum(nurl)
	if n != "" {
		t.Errorf("n=%s should be ''", n)
	}
	nurl = "https://foo.example.com/?page=42"
	n = getPageNum(nurl)
	if n != "42" {
		t.Errorf("n=%s should be 42", n)
	}
	nurl = "https://foo.example.com/?country=fr&page=43"
	n = getPageNum(nurl)
	if n != "43" {
		t.Errorf("n=%s should be 43", n)
	}
	nurl = "https://foo.example.com/?country=fr&page=666&bar=1"
	n = getPageNum(nurl)
	if n != "666" {
		t.Errorf("n=%s should be 666", n)
	}

	nurl = ""
	n = getPageNum(nurl)
	if n != "" {
		t.Errorf("n=%s should be ''", n)
	}
}

func TestClienthandleAPIResponse(t *testing.T) {
	var (
		r       http.Response
		b       bytes.Buffer
		jsonErr string
	)

	client, err := NewClient()
	_, err = client.handleAPIResponse(nil)
	assert.Error(t, err)

	fmt.Fprintf(&b, "%v", jsonErr)

	r = http.Response{StatusCode: 0, Body: ioutil.NopCloser(&b)}
	body, err := client.handleAPIResponse(&r)
	assert.NoError(t, err)
	assert.Empty(t, body)

	r = http.Response{StatusCode: 200, Body: ioutil.NopCloser(&b)}
	body, err = client.handleAPIResponse(&r)
	assert.NoError(t, err)
	assert.Empty(t, body)

	jsonErr = `{"error":{"status": 501, "errors":[{"detail": "test"}],"code": 500, "detail":"issue"}}`

	fmt.Fprintf(&b, "%v", jsonErr)
	r = http.Response{StatusCode: 300, Body: ioutil.NopCloser(&b)}

	body, err = client.handleAPIResponse(&r)
	assert.NoError(t, err)
	assert.Equal(t, jsonErr, string(body))

	fmt.Fprintf(&b, "%v", jsonErr)
	r = http.Response{StatusCode: 500, Body: ioutil.NopCloser(&b)}

	body, err = client.handleAPIResponse(&r)
	assert.Error(t, err)
	assert.Equal(t, jsonErr, string(body))
}

func TestAddQueryParameters(t *testing.T) {
	p := AddQueryParameters("", map[string]string{})
	assert.Equal(t, "", p)
}

func TestAddQueryParameters_1(t *testing.T) {
	p := AddQueryParameters("", map[string]string{"": ""})
	assert.Equal(t, "?=", p)
}

func TestAddQueryParameters_2(t *testing.T) {
	p := AddQueryParameters("foo", map[string]string{"bar": "baz"})
	assert.Equal(t, "foo?bar=baz", p)
}

func TestClient_AddAPIKey(t *testing.T) {
	c, err := NewClient(Config{APIKey: "foo"})
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotEmpty(t, c)

	opts := map[string]string{}

	newk := c.addAPIKey(opts)
	assert.NotEmpty(t, c.config.APIKey)
	assert.Equal(t, 1, len(newk))
	assert.EqualValues(t, map[string]string{"key": "foo"}, newk)
}

func TestClient_PrepareRequest(t *testing.T) {
	c, err := NewClient(Config{endpoint: testURL})
	require.NoError(t, err)

	opts := map[string]string{}
	req := c.prepareRequest("GET", "foo", opts)

	assert.NotNil(t, req)
	assert.IsType(t, (*http.Request)(nil), req)

	res, _ := url.Parse(testURL + "/foo")
	assert.Equal(t, "GET", req.Method)
	assert.EqualValues(t, res, req.URL)
}

func TestClient_PrepareRequest_2(t *testing.T) {
	c, err := NewClient(TesCfg)
	require.NoError(t, err)

	opts := map[string]string{}
	req := c.prepareRequest("GET", "foo", opts)

	assert.NotNil(t, req)
	assert.IsType(t, (*http.Request)(nil), req)

	res, _ := url.Parse(apiEndpoint + "/foo")
	assert.Equal(t, "GET", req.Method)
	assert.EqualValues(t, res, req.URL)
}

func TestClient_PrepareRequest_3(t *testing.T) {
	c, err := NewClient(TesCfg)
	require.NoError(t, err)

	opts := map[string]string{}
	req := c.prepareRequest("FETCH", testURL+"/foo", opts)

	assert.NotNil(t, req)
	assert.IsType(t, (*http.Request)(nil), req)

	res, _ := url.Parse(testURL + "/foo")
	assert.Equal(t, "GET", req.Method)
	assert.EqualValues(t, res, req.URL)
}

func TestClient_MergeGlobalOptions(t *testing.T) {
	c, err := NewClient(TesCfg)
	require.NoError(t, err)

	opts := map[string]string{"foo": "bar"}
	c.opts = map[string]string{"baz": "xyz"}
	res := map[string]string{"foo": "bar", "baz": "xyz"}

	c.mergeGlobalOptions(opts)
	assert.EqualValues(t, res, opts)
}
