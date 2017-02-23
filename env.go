package configpipe

import (
	"os"
	"strings"

	"github.com/mgutz/jo"
)

// Env returns a filter that processes environment variables.
//
// prefix is the prefix of environment variables, defaults to ""
// dotSeparator is the character reprensenting a dot in JSON path. eg dotSeparator="_", user_name becomes key "user.name"
func Env(prefix string, dotSeparator string) Filter {
	L := len(prefix)

	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		o := jo.New()
		for _, item := range os.Environ() {
			splits := strings.Split(item, "=")

			key := splits[0]

			// check prefix and discard it if found, otherwise ignore the key
			if prefix != "" {
				pos := strings.Index(key, prefix)
				if pos != 0 {
					continue
				}
				key = key[L:]
			}

			// allow dot representation, eg "sever_port" => "server.port"
			// some OS do not allow dots in env var name
			if dotSeparator != "" {
				key = strings.Replace(key, dotSeparator, ".", -1)
			}

			val := strings.Join(splits[1:], "=")
			o.Set(key, val)
		}

		m, err := o.Map(".")
		if err != nil {
			return nil, err
		}

		return Merge(m, input)
	})
}
