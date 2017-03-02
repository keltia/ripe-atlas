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
	realPath := path.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", file), "config.toml")
	if res != realPath {
		t.Errorf("Error: badly formed fullname %s—%s", res, realPath)
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

	// Check plain file
	file = "foo.toml"
	res = checkName(file)
	if res != file {
		t.Errorf("Error: plein file.toml is acceptable")
	}
}

func TestLoadConfig(t *testing.T) {
	file := "newconfig.toml"
	conf, err := LoadConfig(file)
	if err != nil {
		t.Errorf("%s does not exist, it should not be an error", file)
	}

	file = "config.toml"
	conf, err = LoadConfig(file)
	if err != nil {
		t.Errorf("Malformed file %s: %v", file, err)
	}

	defaultProbe := 666
	if conf.DefaultProbe != defaultProbe {
		t.Errorf("Malformed default %d: %d\nconf: %#v", conf.DefaultProbe, defaultProbe, conf)
	}

	key := "<INSERT-API-KEY>"
	if conf.APIKey != key {
		t.Errorf("Malformed default %s: %s", conf.APIKey, key)
	}

	poolSize := 10
	if conf.PoolSize != poolSize {
		t.Errorf("Malformed default %d: %d", conf.PoolSize, poolSize)
	}
}
