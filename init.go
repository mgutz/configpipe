package configpipe

import (
	"time"

	"github.com/mgutz/jo"
)

// Configuration wraps a JSON path friendly reader.
type Configuration struct {
	*jo.Object
}

// File is the argument for file-based configurations.
type File struct {
	IgnoreErrors  bool
	Path          string
	Watch         bool
	WatchInterval time.Duration
}

// FilterFunc is a helper that wraps a function to implement
// the Filter interface.
type FilterFunc func(input map[string]interface{}) (map[string]interface{}, error)

// Process implements the Filter interface.
func (ff FilterFunc) Process(input map[string]interface{}) (map[string]interface{}, error) {
	return ff(input)
}

// Filter processes an input map and returns a new map, usually
// merging values over a copy of the input.
type Filter interface {
	Process(input map[string]interface{}) (map[string]interface{}, error)
}
