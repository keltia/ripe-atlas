package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckType(t *testing.T) {
	d := Definition{Type: "foo"}

	valid := checkType(d)
	assert.EqualValues(t, false, valid, "should be false")

	d = Definition{Type: "dns"}
	valid = checkType(d)
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckTypeAs(t *testing.T) {
	d := Definition{Type: "dns"}
	valid := checkTypeAs(d, "foo")
	assert.EqualValues(t, false, valid, "should be false")

	valid = checkTypeAs(d, "dns")
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckAllTypesAs(t *testing.T) {
	dl := []Definition{
		{Type: "foo"},
		{Type: "ping"},
	}

	valid := checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "dns"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "ping"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, true, valid, "should be true")
}

