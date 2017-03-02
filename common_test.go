package atlas

import (
	"testing"
)

func TestSetAuth(t *testing.T) {

	SetAuth("foo")

	if APIKey != "foo" {
		t.Errorf("APIKey is not set")
	}
}

func TestWantAuth(t *testing.T) {

	var (
		key string
		ok  bool
	)

	SetAuth("")
	if key, ok = HasAPIKey(); ok != false {
		t.Errorf("WantAuth() should be true")
	}

	SetAuth("foo")
	if key, ok = HasAPIKey(); ok != true {
		t.Errorf("ok should be true")
	}

	if key != "foo" {
		t.Errorf("APIKey should be %s but is %s", "foo", key)
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
		t.Errorf("n=%s should be ''", n)
	}
}
