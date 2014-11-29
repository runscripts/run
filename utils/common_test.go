package utils

import "testing"

func TestLog(t *testing.T) {
	LogError("Testing LogError...\n")
	LogError("Testing %s...\n", "LogError")
	LogInfo("Testing LogInfo...\n")
	LogInfo("Testing %s...\n", "LogInfo")
}

func TestErrorf(t *testing.T) {
	if Errorf("Testing") == nil || Errorf("Testing %s", "Errorf") == nil {
		t.Errorf("Returned error shouldn't be nil")
	}
}

func TestStrToSha1(t *testing.T) {
	if StrToSha1("Hello") != "f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0" {
		println(StrToSha1("Hello"))
		t.Error("Wrong encryption result")
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists("./common.go") {
		t.Errorf("How does common.go not exist?")
	}
}
