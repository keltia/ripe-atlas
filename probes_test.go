package atlas

import (
	"testing"
)

func TestClient_GetProbe(t *testing.T) {
	/*	defer gock.Off()

		//myurl, _ := url.Parse(apiEndpoint)

		gock.New(apiEndpoint).
			Get("/probes/0").
			MatchParam("key", "foobar").
			Reply(403).
			BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

		c := Before(t)

		gock.InterceptClient(c.client)
		defer gock.RestoreClient(c.client)

		myerr := "status: 403 code: 104 - r:The provided API key does not exist\nerrors: []"

		rp, err := c.GetProbe(0)

		t.Logf("rp=%#v", rp)
		assert.Error(t, err)
		assert.Nil(t, rp)
		assert.Equal(t, myerr, err.Error())
	*/
}
