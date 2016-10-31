package configpipe

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// YAMLFile returns a filter that process a YAML file
func YAMLFile(path string) Filter {
	return func(input map[string]interface{}) (map[string]interface{}, error) {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		obj := map[string]interface{}{}
		err = yaml.Unmarshal(content, &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	}
}

// YAMLString returns a filter that process a YAML encoded string.
func YAMLString(content string) Filter {
	return func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(content), &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	}
}
