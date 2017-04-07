# configpipe

__WARNING__: Unstable. Still early commits. API may switch to Filter interface instead
of Filter function.

Simpler configuration using pipe and filters.


## Motivation

1.  Treat all configuration sources as filters that process and merge map of values.

    * CLI args
    * Environment variables
    * YAML file or string
    * JSON file or string
    * UCL file or string

    The advantage of filters is one can add custom filters such as decrypting
    some keys.

2.  Explicit merge order for overriding values.

3.  Run-time or remote changes

## Usage

    import (
        "os"

        conf "github.com/mgutz/configpipe"
    )

    var config *conf.Configuration

    func decryptor(input map[string] interface{}) (map[string]interface{}, error) {
        // ... decrypt some values, add or remove keys
    }

    func init() {
        envmode := os.Getenv("run_env")

        var prodConfig conf.Filter
        if envmode == "production" {
             prodConfig = conf.YAMLFile(&conf.File{Path: "config.yaml", MustExist: true})
        }

        // filters execute left to right, which means later filters merge over
        // earlier filters
        config, err := conf.Processv(
            // read from config.json file (if present)
            conf.JSONFile(&conf.File{Path: "config.json"}),

            // Any nil filter is noop, so this WILL NOT be processed in development mode.
            prodConfig,

            // read from environment variables that have prefix "CFG_" and replace "_" with "." for JSON Path
            conf.Env("CFG_", "_"),

            // read from argv
            confg.Argv(),

            // use custom filter to decrypt encrypted values
            conf.FilterFunc(decryptor),
        )
    }


## Reading values

    // idiomatic, verbose way
    s, err := config.String("USER")
    n, err := config.Int64("nested.key")

    // default value if missing
    s = config.OrString("USER", "peon")
    n = config.OrInt64("nested.key", 100)

    // zero value if missing
    s = config.AsString("USER") // returns "" if key is missing or cannot be coerced

    // panic if key cannot be coerced or is missing
    s = config.MustString("USER")
