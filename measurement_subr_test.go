package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
