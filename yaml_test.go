package configpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLString(t *testing.T) {
	j1 := `
{
	"one": 1
}
`
	y1 := `
one: 100
`
	config, err := Runv(JSONString(j1), YAMLString(y1))
	assert.NoError(t, err)
	assert.Equal(t, config.MustInt("one"), 100)
}
