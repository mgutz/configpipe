package configpipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ucl1 = `
int = 10
string = "ucl"
two = 1000
`

func TestUCLString(t *testing.T) {
	j := `
two = 2
`
	config, err := Process(UCLString(ucl1), UCLString(j))
	assert.NoError(t, err)
	assert.Equal(t, 10, config.MustInt("int"))
	assert.Equal(t, 2, config.MustInt("two"))
	assert.Equal(t, "ucl", config.MustString("string"))
}

func TestUCLFile(t *testing.T) {
	config, err := Process(
		UCL(&File{Path: "./_fixtures/config.ucl"}),
		UCLString(ucl1),
	)
	assert.NoError(t, err)
	assert.Equal(t, "ucl", config.MustString("format"))
	assert.Equal(t, 91, config.MustInt("numbers[0]"))
}

func TestUCLFileMissing(t *testing.T) {
	_, err := Process(
		UCL(&File{Path: "./_fixtures/missing.ucl"}),
		UCLString(hcl1),
	)
	assert.Error(t, err)
}

func TestUCLFileMissingIgnore(t *testing.T) {
	config, err := Process(
		UCL(&File{Path: "./_fixtures/missing.ucl", IgnoreErrors: true}),
		UCLString(ucl1),
	)
	assert.NoError(t, err)
	assert.Equal(t, 10, config.MustInt("int"))
}

func TestUCLArrayNesting(t *testing.T) {
	const o = `
task {
	name = "build"
	decc = "builds the app"
	bash = <<EOS
echo
EOS
	args {
		chdir = "foobar"
		env {
			prefix = "/usr/local/bin"
		}
	}
}

task {
	name = "clean"
	desc = "removes build files"
	posh = <<EOS
echo
EOS
}
	`
	config, err := Process(
		UCLString(o),
		Trace(),
	)
	assert.NoError(t, err)
	assert.Equal(t, "/usr/local/bin", config.MustString("task[0].args.env.prefix"))
	assert.Equal(t, "build", config.MustString("task[0].name"))
	assert.Equal(t, "clean", config.MustString("task[1].name"))
}
