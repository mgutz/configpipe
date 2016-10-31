package configpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	j1 := `
{
	"one": 1
}
`
	j2 := `
{
	"two": 2
}
`
	config, err := Runv(JSONString(j1), JSONString(j2))
	assert.NoError(t, err)
	assert.Equal(t, config.MustInt("one"), 1)
	assert.Equal(t, config.MustInt("two"), 2)
}
