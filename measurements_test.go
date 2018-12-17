package atlas

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestCheckType(t *testing.T) {
	d := Definition{Type: "foo"}

	valid := checkType(d)
	assert.EqualValues(t, false, valid, "should be false")

	d = Definition{Type: "dns"}
	valid = checkType(d)
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckTypeAs(t *testing.T) {
	d := Definition{Type: "dns"}
	valid := checkTypeAs(d, "foo")
	assert.EqualValues(t, false, valid, "should be false")

	valid = checkTypeAs(d, "dns")
	assert.EqualValues(t, true, valid, "should be true")
}

func TestCheckAllTypesAs(t *testing.T) {
	dl := []Definition{
		{Type: "foo"},
		{Type: "ping"},
	}

	valid := checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "dns"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, false, valid, "should be false")

	dl = []Definition{
		{Type: "ping"},
		{Type: "ping"},
	}
	valid = checkAllTypesAs(dl, "ping")
	assert.EqualValues(t, true, valid, "should be true")
}

func TestClient_GetMeasurement(t *testing.T) {

}

func TestClient_DeleteMeasurement_Nokey(t *testing.T) {
	defer gock.Off()

	pkNumber := 666
	myurl, _ := url.Parse(apiEndpoint)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	gock.New(apiEndpoint).
		Delete(fmt.Sprintf("measurements/%d/", pkNumber)).
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(403).
		BodyString(myrp)

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	err := c.DeleteMeasurement(pkNumber)
	assert.Error(t, err)
}

func TestClient_DeleteMeasurement_Ok(t *testing.T) {
	defer gock.Off()

	pkNumber := 666
	myurl, _ := url.Parse(apiEndpoint)

	myrp := Measurement{ID: 666}
	jrp, err := json.Marshal(myrp)

	gock.New(apiEndpoint).
		Delete(fmt.Sprintf("measurements/%d/", pkNumber)).
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(200).
		BodyString(string(jrp))

	c := Before(t)

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	err = c.DeleteMeasurement(pkNumber)
	assert.NoError(t, err)
}

func TestClient_GetMeasurement_Nokey(t *testing.T) {
	defer gock.Off()

	pkNumber := 666
	myurl, _ := url.Parse(apiEndpoint)

	myrp := `{"error":{"status":403,"code":104,"detail":"The provided API key does not exist","title":"Forbidden"}}`

	gock.New(apiEndpoint).
		Get(fmt.Sprintf("measurements/%d/", pkNumber)).
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(403).
		BodyString(myrp)

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetMeasurement(pkNumber)
	assert.Error(t, err)
	assert.Empty(t, rp)
}

func TestClient_GetMeasurement_Ok(t *testing.T) {
	defer gock.Off()

	pkNumber := 666
	myurl, _ := url.Parse(apiEndpoint)

	myrp := Measurement{ID: 666}
	jrp, err := json.Marshal(myrp)

	gock.New(apiEndpoint).
		Get(fmt.Sprintf("measurements/%d/", pkNumber)).
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(200).
		BodyString(string(jrp))

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	rp, err := c.GetMeasurement(pkNumber)
	assert.NoError(t, err)
	assert.NotEmpty(t, rp)
	assert.Equal(t, pkNumber, rp.ID)
}
