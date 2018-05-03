package atlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestForm = map[string]string{
	"Type":        "foo",
	"AF":          "1",
	"InWifiGroup": "true",
	"Spread":      "1",
}

func TestFillDefinition(t *testing.T) {
	err := FillDefinition(nil, TestForm)

	assert.NoError(t, err)
}

func TestFillDefinition2(t *testing.T) {
	d := &Definition{}
	err := FillDefinition(d, TestForm)

	assert.NoError(t, err)
	assert.Equal(t, "foo", d.Type)
	assert.Equal(t, 1, d.AF)
	assert.True(t, d.InWifiGroup)
}
