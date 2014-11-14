package utils

import "github.com/kylelemons/go-gypsy/yaml"

const (
	CONFIG_PATH = "/etc/runscripts.yml"
	DATA_DIR    = "/var/lib/runscripts"
)

type Config struct {
	Sources       map[string]string
	CacheEnabled  bool
}

// Refer to <http://sweetohm.net/html/go-yaml-parsers.en.html>
func NewConfig(path ...string) *Config {
	file, err := yaml.ReadFile(CONFIG_PATH)
	// NewConfig(path) would only be called in testing.
	if len(path) > 0 {
		file, err = yaml.ReadFile(path[0])
	}
	if err != nil {
		panic(err)
	}

	yamlRoot := toYamlMap(file.Root)
	config := Config{}
	// read "sources"
	sources := toYamlMap(yamlRoot["sources"])
	config.Sources = make(map[string]string)
	for scope, url := range sources {
		config.Sources[scope] = url.(yaml.Scalar).String()
	}
	// read "cache-enabled"
	cacheEnabled := yamlRoot["cache-enabled"].(yaml.Scalar).String()
	if cacheEnabled == "True" {
		config.CacheEnabled = true
	} else if cacheEnabled == "False" {
		config.CacheEnabled = false
	} else {
		panic(LogError("cache-enabled is neither True nor False"))
	}

	return &config
}

func toYamlList(node yaml.Node) (yaml.List) {
	result, ok := node.(yaml.List)
	if !ok {
		panic(LogError("%v is not of type list", node))
	}
	return result
}

func toYamlMap(node yaml.Node) (yaml.Map) {
	result, ok := node.(yaml.Map)
	if !ok {
		panic(LogError("%v is not of type map", node))
	}
	return result
}
