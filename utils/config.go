package utils

import "github.com/kylelemons/go-gypsy/yaml"

const CONFIG_PATH = "/etc/runscripts.yml"

type Config struct {
	Sources         []string
	CacheEnabled    bool
}

func NewConfig(path ...string) Config {
	file, err := yaml.ReadFile(CONFIG_PATH)
	if len(path) > 0 {
		file, err = yaml.ReadFile(path[0])
	}
	if err != nil {
		panic(err)
	}

	yamlRoot := toYamlMap(file.Root)
	config := Config{}

	// read "sources"
	sources := toYamlList(yamlRoot["sources"])
	length := len(sources)
	config.Sources = make([]string, length)
	for i := 0; i < length; i++ {
		config.Sources[i] = sources[i].(yaml.Scalar).String()
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

	return config
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
