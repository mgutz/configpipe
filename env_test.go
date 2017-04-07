package configpipe

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	j1 := `
{
	"one": 1
}
`

	os.Setenv("PREFIX_users_mario", "me")

	config, err := Process(JSONString(j1), Env("PREFIX_", "_"))
	assert.NoError(t, err)
	assert.Equal(t, config.MustInt("one"), 1)
	assert.Equal(t, config.MustString("users.mario"), "me")
	// should not include other vars like USER  due to prefix being specified
	assert.Equal(t, config.AsString("USER"), "")
}
