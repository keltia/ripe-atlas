package main

import (
	"testing"
)

func TestValidateFormat(t *testing.T) {
	err := validateFormat("foo")
	if err == true {
		t.Error("Error: should be false")
	}

	err = validateFormat("xml")
	if err == false {
		t.Error("Error: should be true")
	}
}

func TestAnalyzeTarget(t *testing.T) {
	proto, site, path, port := analyzeTarget("https://atlas.ripe.net/v2/api/")
	if proto != "https" {
		t.Errorf("Error: proto should be https")
	}
	if site != "atlas.ripe.net" {
		t.Errorf("Error: site=%s should be atlas.ripe.net", site)
	}
	if port != 443 {
		t.Errorf("Error: port=%d should be 443", port)
	}
	if path != "/v2/api/" {
		t.Errorf("Error: path=%s should be /v2/api/", path)
	}

	proto, site, path, port = analyzeTarget("http://b2b.cfmu:16443/api/")
	if proto != "http" {
		t.Errorf("Error: proto should be http")
	}
	if site != "b2b.cfmu" {
		t.Errorf("Error: site=%s should be b2b.cfmu", site)
	}
	if port != 16443 {
		t.Errorf("Error: port=%d should be 16443", port)
	}
	if path != "/api/" {
		t.Errorf("Error: path=%s should be /api/", path)
	}

	proto, site, path, port = analyzeTarget("https://www.keltia.net")
	if proto != "https" {
		t.Errorf("Error: proto should be https")
	}
	if site != "www.keltia.net" {
		t.Errorf("Error: site=%s should be www.keltia.net", site)
	}
	if port != 443 {
		t.Errorf("Error: port=%d should be 443", port)
	}
	if path != "/" {
		t.Errorf("Error: path=%s should be /", path)
	}

}
