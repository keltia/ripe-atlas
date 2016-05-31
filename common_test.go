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