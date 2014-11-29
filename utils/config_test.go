package utils

import "testing"

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("../run.yml")
	if err != nil {
		t.Error(err)
	}
	if config.Sources["http"] != "http:%s" {
		t.Errorf("Parse result is not correct")
	}
}
