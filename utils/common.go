package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// Default configuration settings.
var CONFIG_PATH = ""
var DATA_DIR = ""

const (
	MASTER_URL = "https://raw.githubusercontent.com/runscripts/run/master/"
)

// Determine the CONFIG_PATH according to OS type.
func SetConfigPath() {
	if runtime.GOOS == "darwin" {
		CONFIG_PATH = "/usr/local/etc/run.conf"
	} else {
		CONFIG_PATH = "/etc/run.conf"
	}
}

func SetDataDir() {
	if runtime.GOOS == "darwin" {
		DATA_DIR = "/usr/local/var/run"
	} else {
		DATA_DIR = "/usr/local/run"
	}
}

// Determine if run is installed.
func IsRunInstalled() bool {
	return FileExists(CONFIG_PATH) && FileExists(DATA_DIR)
}

// Log error message.
func LogError(format string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, format, args...)
	} else {
		fmt.Fprintf(os.Stderr, format)
	}
}

// Log info message.
func LogInfo(format string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(format, args...)
	} else {
		fmt.Printf(format)
	}
}

// Return formatted error.
func Errorf(format string, args ...interface{}) error {
	if len(args) > 0 {
		return fmt.Errorf(format, args...)
	} else {
		return fmt.Errorf(format)
	}
}

// Print error and exit program.
func ExitError(err error) {
	LogError("%v\n", err)
	os.Exit(1)
}

// Convert string into hash string.
func StrToSha1(str string) string {
	sum := [20]byte(sha1.Sum([]byte(str)))
	return hex.EncodeToString(sum[:])
}

// Execute the command to replace current process.
func Exec(args []string) {
	path, err := exec.LookPath(args[0])
	if err != nil {
		ExitError(err)
	}
	err = syscall.Exec(path, args, os.Environ())
	if err != nil {
		ExitError(err)
	}
}

// Determine if the file exists.
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
