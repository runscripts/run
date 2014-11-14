package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func LogError(format string, args ...interface{}) error {
	if len(args) > 0 {
		return fmt.Errorf(format, args)
	} else {
		return fmt.Errorf(format)
	}
}

func LogInfo(format string, args ... interface{}) (int, error) {
	if len(args) > 0 {
		return fmt.Printf(format, args)
	} else {
		return fmt.Printf(format)
	}
}

func StrToSha1(str string) string {
	sum := [20]byte(sha1.Sum([]byte(str)))
	return hex.EncodeToString(sum[:])
}
