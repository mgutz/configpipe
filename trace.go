package configpipe

import (
	"fmt"

	"github.com/mgutz/jo"
)

// Trace traces the map in a pipeline.
func Trace() Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := jo.NewFromMap(input)
		fmt.Println("TRACE", obj.Prettify())
		return input, nil
	})
}
