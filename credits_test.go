package atlas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCredits_InvalidKey(t *testing.T) {
	defer gock.Off()

	myurl, _ := url.Parse(apiEndpoint)

	gock.New(apiEndpoint).
		Get("credits").
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetCredits()

	assert.Error(t, err)
	assert.Empty(t, rp)
}

func TestClient_GetCredits(t *testing.T) {
	defer gock.Off()

	ft, err := ioutil.ReadFile("testdata/credits.json")
	assert.NoError(t, err)

	gock.New(apiEndpoint).
		Get("credits").
		MatchParam("key", "foobar").
		Reply(200).
		BodyString(string(ft))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetCredits()

	var jft Credits

	err = json.Unmarshal(ft, &jft)
	require.NoError(t, err)

	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.EqualValues(t, &jft, rp)
}
