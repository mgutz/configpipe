package configpipe

import (
	"fmt"
	"io"
	"os"
	"strings"

	ucl "github.com/nahanni/go-ucl"
)

// UCL returns a filter that parses a file.
func UCL(file *File) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		f, err := os.Open(file.Path)
		if err != nil {
			if file.IgnoreErrors {
				return input, nil
			}
			return nil, fmt.Errorf("Could not open file %s", file.Path)
		}

		m, err := parseUCL(f)
		f.Close()
		if err != nil {
			if file.IgnoreErrors {
				return input, nil
			}
			return nil, fmt.Errorf("Could not  parse file %s", file.Path)
		}
		return Merge(m, input)
	})
}

func parseUCL(reader io.Reader) (map[string]interface{}, error) {
	parser := ucl.NewParser(reader)
	return parser.Ucl()
}

// UCLString returns a filter that process an HCL encoded string.
func UCLString(content string) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		reader := strings.NewReader(content)
		m, err := parseUCL(reader)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse UCL string")
		}
		return Merge(m, input)
	})
}
