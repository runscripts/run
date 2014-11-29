package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Default configuration settings.
const (
	CONFIG_PATH = "/etc/run.yml"
	DATA_DIR    = "/var/lib/run"

	RUN_YML_URL = "https://raw.githubusercontent.com/runscripts/run/master/run.yml"
)

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
