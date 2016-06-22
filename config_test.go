package atlas

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestCheckName(t *testing.T) {
	os.Setenv("HOME", "/home/foo")

	// Check tag usage
	file := "mytag"
	res := checkName(file)
	real := path.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", file), "config.toml")
	if res != real {
		t.Errorf("Error: badly formed fullname %s—%s", res, real)
	}

	// Check fullname usage
	file = "/nonexistent/foobar.toml"
	res = checkName(file)
	if res != file {
		t.Errorf("Error: badly formed fullname %s", res)
	}

	// Check bad usage
	file = "/toto.yaml"
	res = checkName(file)
	if res != "" {
		t.Errorf("Error: should end with .toml: %s", res)
	}
}

func TestLoadConfig(t *testing.T) {
	file := "config.toml"
	conf, err := LoadConfig(file)
	if err != nil {
		t.Errorf("Malformed file %s: %v", file, err)
	}

	defProbes := 10
	if conf.DefProbes != defProbes {
		t.Errorf("Malformed default %s: %s", conf.DefProbes, defProbes)
	}

	key := "<INSERT-API-KEY>"
	if conf.ApiKey != key {
		t.Errorf("Malformed default %s: %s", conf.ApiKey, key)
	}
}
