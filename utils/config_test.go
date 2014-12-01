package utils

import "testing"

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("../run.conf")
	if err != nil {
		t.Error(err)
	}
	if config.Sources["http"] != "http:%s" {
		t.Errorf("Parse result is not correct")
	}
}

func TestNewConfigFromString(t *testing.T) {
	str := `
[sources]
# ignore
default: default_url
  space: space_url  `
	config, err := NewConfigFromString(str)
	if err != nil {
		t.Error(err)
	}
	if config.Sources["default"] != "default_url" {
		t.Errorf("'default' source is not parsed correctly")
	} else if config.Sources["space"] != "space_url" {
		t.Errorf("'space' source is not parsed correctly")
	} else if _, ok := config.Sources["ignore"]; ok {
		t.Errorf("'ignore' source should be ignored")
	}
}
