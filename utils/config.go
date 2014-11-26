package utils

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

// Default configuration settings.
const (
	CONFIG_PATH = "/etc/run.yml"
	DATA_DIR    = "/usr/local/run"
	RUN_PATH    = "/usr/bin/run"

	RUN_YML_URL = "https://raw.githubusercontent.com/runscripts/run/master/run.yml"
)

// Configuration ojects in the YAML file.
type Config struct {
	Sources map[string]string
	// Future options can be added here.
}

// Read default YAML file to get configuration.
// Refer to <http://sweetohm.net/html/go-yaml-parsers.en.html> for usage.
func NewConfig(path ...string) *Config {
	file, err := yaml.ReadFile(CONFIG_PATH)
	// NewConfig(path) would be only called in testing.
	if len(path) > 0 {
		file, err = yaml.ReadFile(path[0])
	}
	if err != nil {
		LogError("Failed to parse configuration file %s\n", CONFIG_PATH)
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

// Get YAML list.
func toYamlList(node yaml.Node) yaml.List {
	result, ok := node.(yaml.List)
	if !ok {
		panic(Errorf("%v is not of type list", node))
	}
	return result
}

// Get YAML map.
func toYamlMap(node yaml.Node) yaml.Map {
	result, ok := node.(yaml.Map)
	if !ok {
		panic(Errorf("%v is not of type map", node))
	}
	return result
}
