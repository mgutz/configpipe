package configpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	j := `
{
	"two": 2
}
`
	config, err := Runv(JSONString(json1), JSONString(j))
	assert.NoError(t, err)
	assert.Equal(t, 10, config.MustInt("jint"))
	assert.Equal(t, 2, config.MustInt("two"))
}

func TestJSONFile(t *testing.T) {
	config, err := Runv(
		JSONFile(&File{Path: "./_fixtures/config.json"}),
		JSONString(json1),
	)
	assert.NoError(t, err)
	assert.Equal(t, "json", config.MustString("format"))
	assert.Equal(t, 10, config.MustInt("numbers[0]"))
}

func TestJSONFileMissing(t *testing.T) {
	config, err := Runv(
		JSONFile(&File{Path: "./_fixtures/missing.json"}),
		JSONString(json1),
	)
	assert.NoError(t, err)
	assert.Equal(t, 10, config.MustInt("jint"))
}

func TestJSONFileMissingIgnore(t *testing.T) {
	_, err := Runv(
		JSONFile(&File{Path: "./_fixtures/missing.json", MustExist: true}),
		JSONString(json1),
	)
	assert.Error(t, err)
}
