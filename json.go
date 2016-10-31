package configpipe

import (
	"encoding/json"
	"os"
)

// JSONFile returns a filter that process a JSON file
func JSONFile(path string) Filter {
	return func(input map[string]interface{}) (map[string]interface{}, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		decoder := json.NewDecoder(file)
		content := map[string]interface{}{}
		err = decoder.Decode(&content)
		if err != nil {
			return nil, err
		}
		return Merge(content, input)
	}
}

// JSONString returns a filter that process a JSON encoded string.
func JSONString(content string) Filter {
	return func(input map[string]interface{}) (map[string]interface{}, error) {
		obj := map[string]interface{}{}
		err := json.Unmarshal([]byte(content), &obj)
		if err != nil {
			return nil, err
		}
		return Merge(obj, input)
	}
}
