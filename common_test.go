package atlas

import (
	"testing"
	"reflect"
)

func TestSetAuth(t *testing.T) {

	SetAuth("foo", "bar")

	if APIUser != "foo" {
		t.Errorf("APIUser is not set")
	}

	if APIPassword != "bar" {
		t.Errorf("APIPassword is not set")
	}
}

func TestWantAuth(t *testing.T) {

	SetAuth("", "")
	if WantAuth() != nil {
		t.Errorf("WantAuth() should be nil")
	}

	SetAuth("foo", "bar")
	auth := WantAuth()
	if ta := reflect.TypeOf(auth).String(); ta != "*gopencils.BasicAuth" {
		t.Errorf("auth should be of type %s but is %s\n%#v", "*gopencils.BasicAuth", ta)
	}

	if auth.Username != "foo" {
		t.Errorf("auth.Username should be %s but is %s", "foo", auth.Username)
	}

	if auth.Password != "bar" {
		t.Errorf("auth.Password should be %s but is %s", "bar", auth.Password)
	}
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
		t.Errorf("n=%d should be ''", n)
	}
}
