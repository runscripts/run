package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func LogError(format string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, format, args...)
	} else {
		fmt.Fprintf(os.Stderr, format)
	}
}

func LogInfo(format string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Printf(format, args...)
	} else {
		fmt.Printf(format)
	}
}

func Errorf(format string, args ...interface{}) error {
	if len(args) > 0 {
		return fmt.Errorf(format, args...)
	} else {
		return fmt.Errorf(format)
	}
}

func StrToSha1(str string) string {
	sum := [20]byte(sha1.Sum([]byte(str)))
	return hex.EncodeToString(sum[:])
}

func Exec(args []string) error {
	env := os.Environ()
	var path string
	var err error
	if args[0][0] == '/' {
		path = args[0]
	} else {
		path, err = exec.LookPath(args[0])
		if err != nil {
			panic(err)
		}
	}
	return syscall.Exec(path, args, env)
}
