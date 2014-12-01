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
