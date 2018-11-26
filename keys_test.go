package atlas

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetKey_InvalidKey(t *testing.T) {
	defer gock.Off()

	myurl, _ := url.Parse(apiEndpoint)

	gock.New(apiEndpoint).
		Get("keys/foobar-uuid").
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(403).
		BodyString(`{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`)

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	myerr := "GetKey: The provided API key does not exist"

	rp, err := c.GetKey("foobar-uuid")

	t.Logf("rp=%#v\nerr=%v", rp, err)
	assert.Error(t, err)
	assert.Empty(t, rp)
	assert.Equal(t, myerr, err.Error())

}
