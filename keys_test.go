package atlas

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetKey_BadKey(t *testing.T) {
	defer gock.Off()

	opts1 := map[string]string{
		"key":  "foobar",
		"uuid": "blah",
	}

	gock.New(apiEndpoint).
		Get("/keys/" + opts1["uuid"]).
		MatchParams(opts1).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetKey("27768f56-bb86-11e8-b7d0-27cd18d24377")

	assert.Error(t, err)
	assert.Empty(t, rp)

}

func TestClient_GetKey(t *testing.T) {
	defer gock.Off()

	opts1 := map[string]string{
		"key": "foobar",
	}

	fk, err := ioutil.ReadFile("testdata/single-key.json")
	assert.NoError(t, err)

	gock.New(apiEndpoint).
		Get("/keys/" + opts1["uuid"]).
		MatchParams(opts1).
		Reply(200).
		BodyString(string(fk))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	var jfk Key

	err = json.Unmarshal(fk, &jfk)
	require.NoError(t, err)

	rp, err := c.GetKey("27768f56-bb86-11e8-b7d0-27cd18d24377")

	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.EqualValues(t, jfk, rp)
}

func TestClient_GetKeys_Badkey(t *testing.T) {
	defer gock.Off()

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

	ft, err := ioutil.ReadFile("testdata/keys-list.json")
	assert.NoError(t, err)

	fk, err := ioutil.ReadFile("testdata/keys.json")
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

	var jfk []Key

	err = json.Unmarshal(fk, &jfk)
	require.NoError(t, err)

	rp, err := c.GetKeys(opts)
	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.EqualValues(t, jfk, rp)
}
