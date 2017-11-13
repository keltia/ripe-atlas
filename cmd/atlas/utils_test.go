package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateFormat(t *testing.T) {
	err := validateFormat("foo")
	assert.EqualValues(t, false, err, "should be false")

	err = validateFormat("xml")
	assert.EqualValues(t, true, err, "should be true")
}

func TestAnalyzeTarget(t *testing.T) {
	proto, site, path, port := analyzeTarget("https://atlas.ripe.net/v2/api/")
	assert.EqualValues(t, "https", proto, "should be fhttps")
	assert.EqualValues(t, "atlas.ripe.net", site, "should be atlas.ripe.net")
	assert.EqualValues(t, 443, port, "should be 443")
	assert.EqualValues(t, "/v2/api/", path, "Error: path=%s should be /v2/api/")

	proto, site, path, port = analyzeTarget("http://b2b.cfmu:16443/api/")
	assert.EqualValues(t, "http", proto, "should be http")
	assert.EqualValues(t, "b2b.cfmu", site, "should be b2b.cfmu")
	assert.EqualValues(t, 16443, port, "should be 16443")
	assert.EqualValues(t, "/api/", path, "Error: path=%s should be /api/")

	proto, site, path, port = analyzeTarget("https://www.keltia.net")
	assert.EqualValues(t, "https", proto, "should be https")
	assert.EqualValues(t, "www.keltia.net", site, "should be keltia.net")
	assert.EqualValues(t, 443, port, "should be 443")
	assert.EqualValues(t, "/", path, "should be /")
}
