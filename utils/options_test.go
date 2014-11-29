package utils

import (
	"os"
	"testing"
)

func TestIsScopeNameValid(t *testing.T) {
	if IsScopeNameValid("test%1") {
		t.Errorf("Failed to detect invalid characters")
	}
	if IsScopeNameValid("-test") {
		t.Errorf("First character should not be '-'")
	}
	if IsScopeNameValid(".test") {
		t.Errorf("First character should not be '.'")
	}
	if !IsScopeNameValid("_4.t-X") {
		t.Errorf("This scope name should be valid")
	}
}

func TestNewOptions(t *testing.T) {
	config, err := NewConfig("../run.yml")
	if err != nil {
		t.Fatal(err)
	}

	fakeArgs := []string{
		"run", "-c", "-h", "-i", "lua", "-I", "-u", "-v", "-V",
	}
	os.Args = fakeArgs
	options, err := NewOptions(config)
	if err != nil {
		t.Fatal(err)
	}
	if !(options.Clean && options.Help && options.Interpreter == "lua" &&
		options.Init && options.Update && options.View && options.Version) {
		t.Errorf("Incorrect parse of short arguments")
	}

	fakeArgs = []string{
		"run", "--clean", "--help", "--init", "--update",
		"--view", "--version",
	}
	os.Args = fakeArgs
	options, err = NewOptions(config)
	if err != nil {
		t.Fatal(err)
	}
	if !(options.Clean && options.Help && options.Init && options.Update &&
		options.View && options.Version) {
		t.Errorf("Incorrect parse of long arguments")
	}

	fakeArgs = []string{
		"run", "bitbucket:run/scripts/redis/start", "-d", "test.conf",
	}
	os.Args = fakeArgs
	options, err = NewOptions(config)
	if err != nil {
		t.Fatal(err)
	}
	if options.Scope != "bitbucket" {
		t.Errorf("Incorrect scope")
	}
	if options.Fields[0] != "run" || options.Fields[1] != "scripts" ||
		options.Fields[2] != "redis/start" {
		t.Errorf("Incorrect fields")
	}
	if options.Args[0] != "-d" || options.Args[1] != "test.conf" {
		t.Errorf("Incorrect arguments")
	}
	if options.URL != "https://bitbucket.org/run/scripts/raw/master/redis/start" {
		t.Errorf("Incorrect URL")
	}
	if options.CacheID != StrToSha1("run/scripts/redis/start") {
		t.Errorf("Incorrect CacheID")
	}
	if options.Script != "start" {
		t.Errorf("Incorrect script name")
	}
}
