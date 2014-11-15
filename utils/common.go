package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

func LogError(format string, args ...interface{}) {
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, format, args)
	} else {
		fmt.Fprintf(os.Stderr, format)
	}
}

func LogInfo(format string, args ... interface{}) {
	if len(args) > 0 {
		fmt.Printf(format, args)
	} else {
		fmt.Printf(format)
	}
}

func StrToSha1(str string) string {
	sum := [20]byte(sha1.Sum([]byte(str)))
	return hex.EncodeToString(sum[:])
}
