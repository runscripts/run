package utils

import "testing"

func TestFetch(t *testing.T) {
	err := Fetch("https://raw.githubusercontent.com/runscripts/runscripts/master/.gitignore",
		"/tmp/.gitignore")
	if err != nil {
		t.Error(err)
	}
	err = Fetch("https://raw.githubusercontent.com/runscripts/runscripts/master/gitignore",
		"/tmp/gitignore")
	if err == nil {
		t.Errorf("GET nonexistent URI should return error")
	}
}
