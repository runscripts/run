package utils

import "fmt"

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
