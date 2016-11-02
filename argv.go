package configpipe

import "os"

// Argv parses command line args using minimist package.
func Argv() Filter {
	return ArgvKeys("_nonFlags", "_passthroughArgs")
}

// ArgvKeys parses command line args using minimist package.
func ArgvKeys(nonFlagsKey string, passthroughArgsKey string) Filter {
	return FilterFunc(func(input map[string]interface{}) (map[string]interface{}, error) {
		m := parseArgv(os.Args[1:], nonFlagsKey, passthroughArgsKey)
		return Merge(m, input)
	})
}
