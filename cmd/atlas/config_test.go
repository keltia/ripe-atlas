package main

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestCheckName(t *testing.T) {
	basedir = "/home/foo"

	// Check tag usage
	file := "mytag"
	res := checkName(file)
	realPath := path.Join(basedir, file, "config.toml")
	assert.EqualValues(t, realPath, res, "should be equal")

	// Check fullname usage
	file = "/nonexistent/foobar.toml"
	res = checkName(file)
	assert.EqualValues(t, file, res, "should be equal")

	// Check bad usage
	file = "/toto.yaml"
	res = checkName(file)
	assert.EqualValues(t, "", res, "should be equal")

	file = "/toto.toml"
	res = checkName(file)
	assert.EqualValues(t, file, res, "should be equal")

	file = "toto.yaml"
	res = checkName(file)
	assert.EqualValues(t, "/home/foo/toto.yaml/config.toml", res, "should be equal")

	// Check plain file
	file = "foo.toml"
	res = checkName(file)
	assert.EqualValues(t, file, res, "should be equal")
}

func TestLoadConfig(t *testing.T) {
	file := "newconfig.toml"
	conf, err := LoadConfig(file)
	assert.NoError(t, err, "no file is no error")

	file = "/config.yaml"
	conf, err = LoadConfig(file)
	assert.Error(t, err, "error")
	assert.EqualValues(t, &Config{}, conf, "empty")

	file = "test/perms.toml"
	_, err = LoadConfig(file)
	assert.Error(t, err, "error")

	file = "test/invalid.toml"
	_, err = LoadConfig(file)
	assert.Error(t, err, "error")

	file = "test/config.toml"
	conf, err = LoadConfig(file)
	assert.NoError(t, err, "no error")

	defaultProbe := 666
	assert.EqualValues(t, defaultProbe, conf.DefaultProbe, "should be equal")

	key := "<INSERT-API-KEY>"
	assert.EqualValues(t, key, conf.APIKey, "should be equal")

	poolSize := 10
	assert.EqualValues(t, poolSize, conf.PoolSize, "should be equal")
}
