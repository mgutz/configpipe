package configpipe

import (
	"github.com/imdario/mergo"
	"github.com/mgutz/jo"
)

// ProcessSlice processes a pipeline of configuration filters.
func ProcessSlice(pipeline []Filter) (*Configuration, error) {
	var err error
	m := map[string]interface{}{}
	for _, filter := range pipeline {
		if filter == nil {
			continue
		}
		m, err = filter.Process(m)
		if err != nil {
			return nil, err
		}
	}
	return &Configuration{jo.NewFromMap(m)}, nil
}

// Process processes a pipeline using variable arguments.
func Process(filters ...Filter) (*Configuration, error) {
	return ProcessSlice(filters)
}

// Merge merges two maps returning a new map.
func Merge(left map[string]interface{}, right map[string]interface{}) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	if left != nil {
		if err := mergo.Merge(&m, left); err != nil {
			return nil, err
		}
	}

	if right != nil {
		if err := mergo.Merge(&m, right); err != nil {
			return nil, err
		}
	}

	return m, nil
}
