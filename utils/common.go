package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

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

// Convert string into hash string.
func StrToSha1(str string) string {
	sum := [20]byte(sha1.Sum([]byte(str)))
	return hex.EncodeToString(sum[:])
}

// Execute the command to replace current process.
func Exec(args []string) error {
	path, err := exec.LookPath(args[0])
	if err != nil {
		panic(err)
	}
	return syscall.Exec(path, args, os.Environ())
}

// Determine if the file exists.
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// Determine if run is installed.
func IsRunInstalled() bool {
	return FileExists(CONFIG_PATH) && FileExists(DATA_DIR) && FileExists(RUN_PATH)
}
