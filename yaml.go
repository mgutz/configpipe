package configpipe

import (
	"io/ioutil"

	"github.com/mgutz/yaml"
)

// YAMLFile returns a filter that process a YAML file
func YAMLFile(file *File) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := map[string]interface{}{}
		content, err := ioutil.ReadFile(file.Path)
		if err != nil {
			if file.MustExist {
				return nil, err
			}
			return Merge(obj, input)
		}

		err = yaml.Unmarshal(content, &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	})
}

// YAMLString returns a filter that process a YAML encoded string.
func YAMLString(content string) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(content), &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	})
}
