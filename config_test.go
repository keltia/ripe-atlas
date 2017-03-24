package atlas

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.EqualValues(t, realPath, res, "should be equal")

	// Check fullname usage
	file = "/nonexistent/foobar.toml"
	res = checkName(file)
	assert.EqualValues(t, realPath, res, "should be equal")

	// Check bad usage
	file = "/toto.yaml"
	res = checkName(file)
	assert.EqualValues(t, "", res, "should be equal")

	// Check plain file
	file = "foo.toml"
	res = checkName(file)
	assert.EqualValues(t, file, res, "should be equal")
}

func TestLoadConfig(t *testing.T) {
	file := "newconfig.toml"
	conf, err := LoadConfig(file)
	assert.NoError(t, err, "no file is no error")

	file = "config.toml"
	conf, err = LoadConfig(file)
	assert.NoError(t, err, "no error")

	defaultProbe := 666
	assert.EqualValues(t, defaultProbe, conf.DefaultProbe, "should be equal")

	key := "<INSERT-API-KEY>"
	assert.EqualValues(t, key, conf.APIKey, "should be equal")

	poolSize := 10
	assert.EqualValues(t, poolSize, conf.PoolSize, "should be equal")
}
