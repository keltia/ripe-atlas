package atlas

import (
	"testing"
)

func TestGetPageNum(t *testing.T) {
	url := "https://foo.example.com/"

	n := getPageNum(url)
	if n != 0 {
		t.Errorf("n=%d should be 0", n)
	}
	url = "https://foo.example.com/?page=42"
	n = getPageNum(url)
	if n != 42 {
		t.Errorf("n=%d should be 42", n)
	}
	url = "https://foo.example.com/?country=fr&page=43"
	n = getPageNum(url)
	if n != 43 {
		t.Errorf("n=%d should be 43", n)
	}
	url = "https://foo.example.com/?country=fr&page=666&bar=1"
	n = getPageNum(url)
	if n != 666 {
		t.Errorf("n=%d should be 666", n)
	}
}