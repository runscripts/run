package utils

import (
	"runtime"
	"strconv"
	"strings"
)

// Get the version of go.
func GoVersion() string {
	return runtime.Version()
}

// Determine whether the version is 1.3+ to support syscall.FcntlFlock.
func IsSupportFLock() bool {
	versions := strings.Split(GoVersion()[2:], ".")

	majorVersion, _ := strconv.Atoi(versions[0])
	if majorVersion > 1 {
		return true
	} else if majorVersion == 1 {
		minorVersion, _ := strconv.Atoi(versions[1])
		if minorVersion >= 3 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
