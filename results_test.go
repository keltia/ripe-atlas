package atlas

import (
	"net/url"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_FetchResult(t *testing.T) {
	defer gock.Off()

	myurl, _ := url.Parse("https://example.com/fetch")

	gock.New("https://example.com").
		Get("/fetch").
		Reply(200).
		BodyString(`{"some": "results"`)

	c := Before(t)
	c.level = 2
	require.NotNil(t, c)
	require.NotNil(t, c.client)

	opts := make(map[string]string)
	c.mergeGlobalOptions(opts)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	body, err := c.FetchResult(myurl.String())
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	assert.Equal(t, `{"some": "results"`, body)
}

func TestClient_FetchResult2(t *testing.T) {
	defer gock.Off()

	myurl, _ := url.Parse("https://example.com/fetch")

	gock.New("https://example.com").
		Get("/fetch").
		MatchHeader("foo", "bar").
		Reply(200).
		BodyString(`{"some": "results"`)

	c := Before(t)
	c.level = 2
	require.NotNil(t, c)
	require.NotNil(t, c.client)

	opts := make(map[string]string)
	c.mergeGlobalOptions(opts)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	body, err := c.FetchResult(myurl.String())
	assert.Error(t, err)
	assert.Empty(t, body)
}
