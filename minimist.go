package configpipe

import (
	"os"
	"regexp"

	"github.com/mgutz/jo"
	//"strings"
)

func nextString(list []string, i int) *string {
	if i+1 < len(list) {
		return &list[i+1]
	}
	return nil
}

func sliceContains(slice []string, needle string) bool {
	if slice == nil {
		return false
	}
	for _, s := range slice {
		if s == needle {
			return true
		}
	}
	return false
}

var integerRe = regexp.MustCompile(`^-?\d+$`)
var numberRe = regexp.MustCompile(`^-?\d+(\.\d+)?(e-?\d+)?$`)

// --port=8000
var longFormEqualRe = regexp.MustCompile(`^--.+=`)
var longFormEqualValsRe = regexp.MustCompile(`^--([^=]+)=(.*)$`)

// --port 8000
var longFormRe = regexp.MustCompile(`^--.+`)
var longFormKeyRe = regexp.MustCompile(`^--(.+)`)

//longFormSpaceValsRe := regexp.MustCompile(`^--([^=])=([\s\S]*)$`)

// --no-debug
var negateRe = regexp.MustCompile(`^--no-.+`)
var negateValsRe = regexp.MustCompile(`^--no-(.+)`)

// -abc
var shortFormRe = regexp.MustCompile(`^-[^-]+`)

var lettersRe = regexp.MustCompile(`^[A-Za-z]`)

var notWordRe = regexp.MustCompile(`\W`)

var dashesRe = regexp.MustCompile(`^(-|--)`)

var trueFalseRe = regexp.MustCompile(`^(true|false)`)

// ParseArgs parses os.Args excluding os.Args[0].
func parseArgs() map[string]interface{} {
	return parseArgv(os.Args[1:], "_nonFlags", "_passthroughArgs")
}

// ParseArgv parses an argv for options.
//
// nonFlagsKey is the key for any option that is not a leading dash. For example `app --foo fruit -b apple`,
//		nonFlags == []string{"apple"}
// passthroughKey is the key for arguments that are pass through, For example `app -foo bar -- --foo=bar --bar=bah
//		passthrough == []string{"--foo=bar", "--bar=bah"}
func parseArgv(argv []string, nonFlagsKey string, passthroughKey string) map[string]interface{} {
	rest := []string{}
	result := jo.New()

	setKV := func(key string, val interface{}) {
		result.Set(key, val)
	}

	l := len(argv)
	argsAt := func(i int) string {
		if i > -1 && i < l {
			return argv[i]
		}
		return ""
	}

	i := 0
	for i < len(argv) {
		arg := argv[i]

		if arg == "--" {
			setKV(passthroughKey, argv[i+1:])
			break
		}

		argAt := func(i int) string {
			if i >= 0 && i < len(arg) {
				return arg[i : i+1]
			}
			return ""
		}
		if longFormEqualRe.MatchString(arg) {
			// --long-form=value

			m := longFormEqualValsRe.FindStringSubmatch(arg)
			//fmt.Printf("--long-form= %s\n", arg)
			setKV(m[1], m[2])

		} else if negateRe.MatchString(arg) {
			//fmt.Printf("--no-flag %s\n", arg)

			m := negateValsRe.FindStringSubmatch(arg)
			setKV(m[1], false)

		} else if longFormRe.MatchString(arg) {
			// --long-form
			//fmt.Printf("--long-form %s\n", arg)

			key := longFormKeyRe.FindStringSubmatch(arg)[1]
			next := argsAt(i + 1)

			if next == "" {
				// --arg
				setKV(key, true)
			} else if next[0:1] == nonFlagsKey {
				// --arg -o | --arg --other
				setKV(key, true)
			} else {
				setKV(key, next)
				i++
			}
		} else if shortFormRe.MatchString(arg) {
			// -abc a, b are boolean c is undetermined
			//fmt.Printf("-short-form %s\n", arg)

			letters := arg[1:]

			L := len(letters)
			lettersAt := func(i int) string {
				if i < L {
					return letters[i : i+1]
				}
				return ""
			}

			broken := false
			k := 0
			for k < len(letters) {
				next := arg[k+2:]
				if next == nonFlagsKey {
					setKV(lettersAt(k), next)
					k++
					continue
				}
				if lettersRe.MatchString(lettersAt(k)) && numberRe.MatchString(next) {
					setKV(lettersAt(k), next)
					broken = true
					break
				}
				if k+1 < len(letters) && notWordRe.MatchString(lettersAt(k+1)) {
					setKV(lettersAt(k), next)
					broken = true
					break
				}

				setKV(lettersAt(k), true)
				k++
			}

			key := argAt(len(arg) - 1)
			if !broken && key != nonFlagsKey {

				if i+1 < len(argv) {
					nextArg := argv[i+1]
					if !dashesRe.MatchString(nextArg) {
						setKV(key, nextArg)
						i++
					}
				} else {
					setKV(key, true)
				}
			}
		} else {
			rest = append(rest, arg)
			setKV(nonFlagsKey, rest)
		}

		i++
	}

	return result.AsMap(".")
}
