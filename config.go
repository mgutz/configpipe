package configpipe

import (
	"github.com/imdario/mergo"
	"github.com/mgutz/jo"
)

// Run runs a pipeline of configuration filters.
func Run(pipeline []Filter) (*Configuration, error) {
	var err error
	m := map[string]interface{}{}
	for _, filter := range pipeline {
		if filter == nil {
			continue
		}
		m, err = filter(m)
		if err != nil {
			return nil, err
		}
	}
	return &Configuration{jo.NewFromMap(m)}, nil
}

// Runv processes a pipeline using variable arguments.
func Runv(filters ...Filter) (*Configuration, error) {
	return Run(filters)
}

// Merge merges two maps returning a new map.
func Merge(left map[string]interface{}, right map[string]interface{}) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	err := mergo.Merge(&m, left)
	if err != nil {
		return nil, err
	}
	err = mergo.Merge(&m, right)
	if err != nil {
		return nil, err
	}
	return m, nil
}
