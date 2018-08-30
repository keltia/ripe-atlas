package atlas

import (
	"testing"

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
