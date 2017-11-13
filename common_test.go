package atlas

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sendgrid/rest"
)

func TestSetAuth(t *testing.T) {

	SetAuth("foo")

	if APIKey != "foo" {
		t.Errorf("APIKey is not set")
	}
}

func TestGetVersion(t *testing.T) {
	ver := GetVersion()
	assert.EqualValues(t, ourVersion, ver, "should be equal")
}

func TestWantAuth(t *testing.T) {

	var (
		key string
		ok  bool
	)

	SetAuth("")
	assert.EqualValues(t, "", APIKey, "should be equal")

	key, ok = HasAPIKey()
	assert.EqualValues(t, false, ok, "should be equal")
	assert.EqualValues(t, "", key, "should be equal")


	SetAuth("foo")
	key, ok = HasAPIKey()
	assert.EqualValues(t, true, ok, "should be equal")
	assert.EqualValues(t, "foo", key, "should be equal")
}

func TestGetPageNum(t *testing.T) {
	url := "https://foo.example.com/"

	n := getPageNum(url)
	if n != "" {
		t.Errorf("n=%s should be ''", n)
	}
	url = "https://foo.example.com/?page=42"
	n = getPageNum(url)
	if n != "42" {
		t.Errorf("n=%s should be 42", n)
	}
	url = "https://foo.example.com/?country=fr&page=43"
	n = getPageNum(url)
	if n != "43" {
		t.Errorf("n=%s should be 43", n)
	}
	url = "https://foo.example.com/?country=fr&page=666&bar=1"
	n = getPageNum(url)
	if n != "666" {
		t.Errorf("n=%s should be 666", n)
	}

	url = ""
	n = getPageNum(url)
	if n != "" {
		t.Errorf("n=%s should be ''", n)
	}
}

func TestHandleAPIResponse(t *testing.T) {
	var r rest.Response

	err := handleAPIResponse(nil)
	assert.Error(t, err, "should be in error")


	r = rest.Response{StatusCode: 0}
	err = handleAPIResponse(&r)
	assert.NoError(t, err, "should be no error")

	r = rest.Response{StatusCode: 200}
	err = handleAPIResponse(&r)
	assert.NoError(t, err, "should be no error")

	var jsonErr = `error:{status: 501, code: 500, detail: "test"}`

	r.StatusCode = 300
	r.Body = jsonErr
	err = handleAPIResponse(&r)
	assert.NoError(t, err, "should be in error")

	r.StatusCode = 500
	r.Body = jsonErr
	err = handleAPIResponse(&r)
	assert.Error(t, err, "should be in error")
}
