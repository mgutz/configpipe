# configpipe

Configuration through pipeline

Simpler configuration using pipe and filters.

## Usage

    import cp "github.com/mgutz/configpipe"

    config, err := cp.Runv(
        // read from environment variables that have prefix "PREFIX_" and replace "_" with "." for JSON Path
        cp.Env("PREFIX_"), "_"),
        // read from config.json file
        cp.JSONFile("config.json"),
    )


    // To read values
    config.String("USER")
    config.Int64("nested.key")



