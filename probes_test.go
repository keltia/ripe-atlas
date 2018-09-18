package atlas

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetProbe_Badkey(t *testing.T) {
	defer gock.Off()

	//myurl, _ := url.Parse(apiEndpoint)

	gock.New(apiEndpoint).
		Get("/probes/0").
		MatchParam("key", "foobar").
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	//myerr := "status: 403 code: 104 - r:The provided API key does not exist\nerrors: []"

	rp, err := c.GetProbe(0)

	t.Logf("rp=%#v", rp)
	assert.Error(t, err)
	assert.Empty(t, rp)
}

func TestClient_GetProbe(t *testing.T) {
	defer gock.Off()

	//myurl, _ := url.Parse(apiEndpoint)

	ft, err := ioutil.ReadFile("testdata/probe-0.json")
	assert.NoError(t, err)

	gock.New(apiEndpoint).
		Get("/probes/0").
		MatchParam("key", "foobar").
		Reply(200).
		BodyString(string(ft))

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetProbe(0)

	var jft Probe

	err = json.Unmarshal(ft, &jft)
	require.NoError(t, err)

	t.Logf("rp=%#v", rp)
	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.EqualValues(t, &jft, rp)
}
