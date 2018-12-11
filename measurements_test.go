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
	defer gock.Off()

	id := 666
	myurl, _ := url.Parse(apiEndpoint)

	myrp := Measurement{}
	jr, _ := json.Marshal(myrp)

	gock.New(apiEndpoint).
		Get("measurements/"+fmt.Sprintf("%d", id)).
		MatchParam("key", "foobar").
		MatchHeaders(map[string]string{
			"host":       myurl.Host,
			"user-agent": fmt.Sprintf("ripe-atlas/%s", ourVersion),
		}).
		Reply(200).
		BodyString(string(jr))

	c := Before(t)
	c.level = 2

	gock.InterceptClient(c.client)
	defer gock.RestoreClient(c.client)

	m, err := c.GetMeasurement(id)
	t.Logf("err=%v", err)
	assert.NoError(t, err)
	assert.EqualValues(t, &myrp, m)
}
