package atlas

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetKeys_Badkey(t *testing.T) {
	defer gock.Off()

	//myurl, _ := url.Parse(apiEndpoint)

	gock.New(apiEndpoint).
		Get("/keys").
		MatchParam("key", "foobar").
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	opts := map[string]string{}
	rp, err := c.GetKeys(opts)

	assert.Error(t, err)
	assert.Empty(t, rp)
}

func TestClient_GetKeys(t *testing.T) {
	defer gock.Off()

	ft, err := ioutil.ReadFile("testdata/keys.json")
	assert.NoError(t, err)

	fk, err := ioutil.ReadFile("testdata/key.json")
	assert.NoError(t, err)

	gock.New(apiEndpoint).
		Get("/keys").
		MatchParam("key", "foobar").
		Reply(200).
		BodyString(string(ft))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	opts := map[string]string{}
	rp, err := c.GetKeys(opts)

	var jfk []Key

	err = json.Unmarshal(fk, &jfk)
	require.NoError(t, err)

	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.EqualValues(t, jfk, rp)
}
