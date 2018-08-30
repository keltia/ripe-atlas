package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeOptions(t *testing.T) {

	var (
		v1, v2, v3 string
	)

	v1 = "aa"
	v2 = "bb"
	v3 = "cc"

	foo := map[string]*string{
		"a": &v1,
		"b": &v2,
	}
	bar := map[string]string{
		"c": v3,
	}

	res := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
	}
	baz := mergeOptions(bar, foo)
	assert.EqualValues(t, res, baz, "equal")
}
