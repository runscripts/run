package utils

import "github.com/kylelemons/go-gypsy/yaml"

const (
	CONFIG_PATH = "/etc/runscripts.yml"
	DATA_DIR    = "/var/lib/runscripts"
)

type Config struct {
	Sources       map[string]string
	// future options can be added here
}

// Refer to <http://sweetohm.net/html/go-yaml-parsers.en.html>
func NewConfig(path ...string) *Config {
	file, err := yaml.ReadFile(CONFIG_PATH)
	// NewConfig(path) would be only called in testing.
	if len(path) > 0 {
		file, err = yaml.ReadFile(path[0])
	}
	if err != nil {
		LogError("failed to parse %s\n", CONFIG_PATH)
		panic(err)
	}

	config := Config{}
	yamlRoot := toYamlMap(file.Root)
	sources := toYamlMap(yamlRoot["sources"])
	config.Sources = make(map[string]string)
	for scope, url := range sources {
		config.Sources[scope] = url.(yaml.Scalar).String()
	}
	return &config
}

func toYamlList(node yaml.Node) (yaml.List) {
	result, ok := node.(yaml.List)
	if !ok {
		panic(Errorf("%v is not of type list", node))
	}
	return result
}

func toYamlMap(node yaml.Node) (yaml.Map) {
	result, ok := node.(yaml.Map)
	if !ok {
		panic(Errorf("%v is not of type map", node))
	}
	return result
}
