package utils

import (
	"os"
	"strings"
)

const VALID_SCOPE_CHARS = "-_." + "1234567890" +
	"abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Determine if a scope name contains invalid characters.
func IsScopeNameValid(scope string) bool {
	for _, c := range scope {
		if strings.IndexRune(VALID_SCOPE_CHARS, c) < 0 {
			return false
		}
	}
	return scope[0] != '-' && scope[0] != '.'
}

// All the options when run the script.
type Options struct {
	Program     string
	Clean       bool
	Help        bool
	Interpreter string
	Init        bool
	Update      bool
	View        bool
	Version     bool
	Scope       string
	Fields      []string
	Args        []string
	URL         string
	CacheID     string
	Script      string
}

// Give configuration and set runtime options.
func NewOptions(config *Config) (*Options, error) {
	options := Options{
		Program:     os.Args[0],
		Clean:       false,
		Help:        false,
		Interpreter: "",
		Init:        false,
		Update:      false,
		View:        false,
		Version:     false,
		Scope:       "default",
		Fields:      nil,
		Args:        []string{},
	}

	// Parse command parameters to options.
	i := 1
	length := len(os.Args)
	for ; i < length; i++ {
		opt := os.Args[i]
		if opt[0] != '-' {
			break
		}
		switch opt {
		case "-c", "--clean":
			options.Clean = true
		case "-h", "--help":
			options.Help = true
		case "-i":
			if i+1 == length || os.Args[i+1][0] == '-' {
				return nil, Errorf("Missing interpreter (e.g., bash, python) after -i")
			} else {
				options.Interpreter = os.Args[i+1]
				i += 1
			}
		case "-I", "--init":
			options.Init = true
		case "-u", "--update":
			options.Update = true
		case "-v", "--view":
			options.View = true
		case "-V", "--version":
			options.Version = true
		default:
			return nil, Errorf("Unknown option %s", opt)
		}
	}

	// Parse full script name to options.
	if i < length {
		// Parse scope.
		var fields string
		opt := os.Args[i]
		j := strings.Index(opt, ":")
		if j >= 0 {
			options.Scope = opt[:j]
			if !IsScopeNameValid(options.Scope) {
				return nil, Errorf("Scope name %s is invalid", options.Scope)
			}
			fields = opt[j+1:]
		} else {
			fields = opt
		}
		// Get scope pattern from configuration.
		pattern, ok := (*config).Sources[options.Scope]
		if !ok {
			return nil, Errorf(
				"%s: sources: %s scope didn't exist",
				CONFIG_PATH, options.Scope,
			)
		}
		// Examine whether scope pattern and fields can be matched.
		nReplace := strings.Count(pattern, "%s")
		if nReplace == 0 {
			return nil, Errorf(
				"%s: sources: %s scope didn't contain %%s",
				CONFIG_PATH, options.Scope,
			)
		} else if nReplace > strings.Count(fields, "/")+1 {
			return nil, Errorf(
				"%s: sources: %s scope required more fields",
				CONFIG_PATH, options.Scope,
			)
		}
		// Parse fields and url.
		options.Fields = strings.SplitN(fields, "/", nReplace)
		for _, f := range options.Fields {
			pattern = strings.Replace(pattern, "%s", f, 1)
		}
		options.URL = pattern
		// Parse args of the script.
		if i+1 < length {
			options.Args = os.Args[i+1:]
		}
		// Parse script name and its cache id.
		options.Script = fields[strings.LastIndex(fields, "/")+1:]
		options.CacheID = StrToSha1(fields)
	}

	return &options, nil
}
