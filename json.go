package configpipe

import (
	"encoding/json"
	"os"
)

// JSON returns a filter that process a JSON file
func JSON(file *File) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		content := map[string]interface{}{}
		f, err := os.Open(file.Path)
		if err != nil {
			if file.IgnoreErrors {
				return Merge(content, input)
			}
			return nil, err
		}

		decoder := json.NewDecoder(f)
		err = decoder.Decode(&content)
		if err != nil {
			return nil, err
		}
		return Merge(content, input)
	})
}

// JSONString returns a filter that process a JSON encoded string.
func JSONString(content string) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := map[string]interface{}{}
		err := json.Unmarshal([]byte(content), &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	})
}
