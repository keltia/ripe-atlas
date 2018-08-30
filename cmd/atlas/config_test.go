package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckName(t *testing.T) {
	basedir = "/home/foo"

	// Check fullname usage
	file := "/nonexistent/foobar.toml"
	res := checkName(file)
	assert.Equal(t, file, res)

	// Check bad usage
	file = "/toto.yaml"
	res = checkName(file)
	assert.Empty(t, res)

	file = "/toto.toml"
	res = checkName(file)
	assert.Equal(t, file, res)

	file = "toto.yaml"
	res = checkName(file)
	assert.Empty(t, res)

	// Check plain file
	file = "foo.toml"
	res = checkName(file)
	assert.EqualValues(t, file, res)
}

func TestLoadConfig(t *testing.T) {
	file := "newconfig.toml"
	conf, err := LoadConfig(file)
	assert.NoError(t, err, "no file is no error")

	file = "/config.yaml"
	conf, err = LoadConfig(file)
	assert.Error(t, err, "error")
	assert.Empty(t, conf)

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
	assert.Equal(t, defaultProbe, conf.DefaultProbe, "should be equal")

	key := "<INSERT-API-KEY>"
	assert.Equal(t, key, conf.APIKey, "should be equal")

	poolSize := 10
	assert.Equal(t, poolSize, conf.ProbeSet.PoolSize, "should be equal")
}
