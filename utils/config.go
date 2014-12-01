package utils

import (
	"bufio"
	"io/ioutil"
	"strings"
)

// Configuration object.
type Config struct {
	Sources map[string]string
	// Future options can be added here.
}

// Read the configuration on default or specified path.
func NewConfig(path ...string) (*Config, error) {
	var content []byte
	var err error
	// NewConfig(path) would be only called in testing.
	if len(path) > 0 {
		content, err = ioutil.ReadFile(path[0])
	} else {
		content, err = ioutil.ReadFile(CONFIG_PATH)
	}
	if err != nil {
		return nil, err
	}

	return NewConfigFromString(string(content))
}

// Parse the string to Config object.
func NewConfigFromString(str string) (*Config, error) {
	config := Config{}
	config.Sources = make(map[string]string)

	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	option := config.Sources
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		} else if line[0] == '[' {
			switch line {
			case "[sources]": option = config.Sources
			default: return nil, Errorf("Unknown option: %s", line)
			}
		} else {
			tokens := strings.SplitN(line, ":", 2)
			if len(tokens) != 2 {
				return nil, Errorf("Incorret configuration format: %s", line)
			}
			option[strings.TrimSpace(tokens[0])] = strings.TrimSpace(tokens[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}
