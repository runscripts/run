package utils

import (
	"os"
	"strings"
)

// All the options when run the script.
type Options struct {
	Program     string
	Help        bool
	Interpreter string
	Update      bool
	View        bool
	Clean       bool
	Version     bool
	Scope       string
	Fields      []string
	Args        []string
	URL         string
	CacheID     string
	Script      string
}

// Give configuration and set runtime options.
func NewOptions(config *Config) *Options {
	options := Options{
		Program:     os.Args[0],
		Help:        false,
		Interpreter: "",
		Update:      false,
		View:        false,
		Clean:       false,
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
		case "-h", "--help":
			options.Help = true
		case "-i":
			if i+1 == length || os.Args[i+1][0] == '-' {
				panic(Errorf("Missing interpreter (e.g., bash, python) after -i"))
			} else {
				options.Interpreter = os.Args[i+1]
				i += 1
			}
		case "-u", "--update":
			options.Update = true
		case "-v", "--view":
			options.View = true
		case "-V", "--version":
			options.Version = true
		case "-c", "--clean":
			options.Clean = true
		default:
			panic(Errorf("Unknown option %s", opt))
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
			fields = opt[j+1:]
		} else {
			fields = opt
		}
		// Get scope pattern from configuration.
		pattern, ok := (*config).Sources[options.Scope]
		if !ok {
			panic(Errorf(
				"%s: sources: %s scope didn't exist",
				CONFIG_PATH, options.Scope,
			))
		}
		// Examine whether scope pattern and fields can be matched.
		nReplace := strings.Count(pattern, "%s")
		if nReplace == 0 {
			panic(Errorf(
				"%s: sources: %s scope didn't contain %%s",
				CONFIG_PATH, options.Scope,
			))
		} else if nReplace > strings.Count(fields, "/")+1 {
			panic(Errorf(
				"%s: sources: %s scope required more fields",
				CONFIG_PATH, options.Scope,
			))
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

	return &options
}
