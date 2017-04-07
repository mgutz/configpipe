package configpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLString(t *testing.T) {
	j := `
two: 2
`
	config, err := Process(YAMLString(yaml1), YAMLString(j))
	assert.NoError(t, err)
	assert.Equal(t, 100, config.MustInt("yint"))
	assert.Equal(t, 2, config.MustInt("two"))
}

func TestYAMLFile(t *testing.T) {
	config, err := Process(
		YAML(&File{Path: "./_fixtures/config.yaml"}),
		YAMLString(yaml1),
	)
	assert.NoError(t, err)
	assert.Equal(t, "yaml", config.MustString("format"))
	assert.Equal(t, 101, config.MustInt("numbers[0]"))
}

func TestYAMLFileMissing(t *testing.T) {
	_, err := Process(
		YAML(&File{Path: "./_fixtures/missing.yaml"}),
		YAMLString(yaml1),
	)
	assert.Error(t, err)
}

func TestYAMLFileMissingIgnore(t *testing.T) {
	config, err := Process(
		YAML(&File{Path: "./_fixtures/missing.yaml", IgnoreErrors: true}),
		YAMLString(yaml1),
	)
	assert.NoError(t, err)
	assert.Equal(t, 100, config.MustInt("yint"))
}
