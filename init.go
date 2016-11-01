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
	MustExist     bool
	Path          string
	Watch         bool
	WatchInterval time.Duration
}

// Filter is a step in the configuration pipeline.
type Filter func(input map[string]interface{}) (map[string]interface{}, error)
