package utils

import (
	"os"
	"strings"
)

type Options struct {
	Program  string
	Help     bool
	Intprt   string
	Update   bool
	View     bool
	Clean    bool
	Version  bool
	Scope    string
	Fields   []string
	Args     string
	URL      string
	CacheID  string
}

func NewOptions(config *Config) *Options {
	options := Options{
		Program : os.Args[0],
		Help    : false,
		Intprt  : "",
		Update  : false,
		View    : false,
		Clean   : false,
		Version : false,
		Scope   : "default",
		Fields  : nil,
		Args    : "",
	}

	// parse options
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
			if i + 1 == length || os.Args[i+1][0] == '-' {
				panic(Errorf("missing interpreter (e.g., bash, python) after -i"))
			} else {
				options.Intprt = os.Args[i+1]
				i += 1
			}
		case "-u", "--update":
			options.Update = true
		case "-v", "--view":
			options.View = true
		case "--clean":
			options.Clean = true
		case "--version":
			options.Version = true
		default:
			panic(Errorf("unknown option %s", opt))
		}
	}

	if i < length {
		// options.Scope
		var fields string
		opt := os.Args[i]
		j := strings.Index(opt, ":")
		if j >= 0 {
			options.Scope = opt[:j]
			fields = opt[j+1:]
		} else {
			fields = opt
		}
		// get scope pattern from config
		pattern, ok := (*config).Sources[options.Scope]
		if !ok {
			panic(Errorf(
				"%s: sources: %s scope didn't exist",
				CONFIG_PATH, options.Scope,
			))
		}
		// examine whether scope pattern and fields can be matched
		nReplace := strings.Count(pattern, "%s")
		if nReplace == 0 {
			panic(Errorf(
				"%s: sources: %s scope didn't contain %%s",
				CONFIG_PATH, options.Scope,
			))
		} else if nReplace > strings.Count(fields, "/") + 1 {
			panic(Errorf(
				"%s: sources: %s scope required more fields",
				CONFIG_PATH, options.Scope,
			))
		}
		// options.Fields
		options.Fields = strings.SplitN(fields, "/", nReplace)
		for _, f := range(options.Fields) {
			pattern = strings.Replace(pattern, "%s", f, 1)
		}
		// options.URL
		options.URL = pattern
		// options.Args
		if i + 1 < length {
			options.Args = strings.Join(os.Args[i+1:], " ")
		}
		// options.CacheID
		scriptFileName := fields[strings.LastIndex(fields, "/") + 1:]
		options.CacheID = scriptFileName + "-" + StrToSha1(fields)
	}

	return &options
}
