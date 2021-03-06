package configpipe

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple(t *testing.T) {
	os.Setenv("PREFIX_users_mario", "me")

	config, err := Process(
		JSONString(json1),
		Trace(),
		JSON(&File{Path: "./_fixtures/config.json"}),
		Trace(),
		YAML(&File{Path: "./_fixtures/config.yaml"}),
		Trace(),
		Env("PREFIX_", "_"),
		Trace(),
		UCLString(ucl1),
		Trace(),
		UCL(&File{Path: "./_fixtures/config.ucl"}),
		Trace(),
	)

	assert.NoError(t, err)
	assert.Equal(t, config.MustInt("jint"), 10)
	assert.Equal(t, config.MustString("users.mario"), "me")
	assert.Equal(t, config.MustString("json.key"), "jsonstring")
	assert.Equal(t, config.MustString("yaml.key"), "yamlstring")
	assert.Equal(t, config.MustString("ucl.key"), "uclstring")

	// should not include other vars like USER  due to prefix being specified
	assert.Equal(t, config.AsString("USER"), "")
}
