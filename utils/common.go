package utils

import "fmt"

func LogError(format string, args ...interface{}) error {
	return fmt.Errorf(format, args)
}

func LogInfo(format string, args ... interface{}) (int, error) {
	return fmt.Printf(format, args)
}
