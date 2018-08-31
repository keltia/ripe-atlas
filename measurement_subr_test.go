package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TesCfg = Config{
		APIKey:       "",
		DefaultProbe: 666,
		IsOneOff:     true,
		PoolSize:     10,
		AreaType:     "country",
		AreaValue:    "fr",
		Tags:         "",
		Verbose:      false,
		Log:          nil,
	}
)

func Before(t *testing.T) *Client {

	testc, err := NewClient(TesCfg)

	assert.NoError(t, err)
	assert.NotNil(t, testc)
	assert.IsType(t, (*Client)(nil), testc)

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
