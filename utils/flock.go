package utils

import (
	"os"
	"syscall"
)

var lockFile *os.File

// Lock the file.
func Flock(path string) error {
	err := fcntlFlock(syscall.F_WRLCK, path)
	if err != nil {
		LogError("Failed to lock file %s\n", path)
		panic(err)
	} else {
		return nil
	}
}

// Unlock the file.
func Funlock(path string) error {
	err := fcntlFlock(syscall.F_UNLCK)
	if err != nil {
		LogError("Failed to unlock file %s\n", path)
		panic(err)
	} else {
		return lockFile.Close()
	}
}

// Control the lock of file.
func fcntlFlock(lockType int16, path ...string) error {
	var err error
	if lockType != syscall.F_UNLCK {
		mode := syscall.O_CREAT | syscall.O_WRONLY
		lockFile, err = os.OpenFile(path[0], mode, 0666)
		if err != nil {
			LogError("Can't open the lock file %s\n", path[0])
			panic(err)
		}
	}

	lock := syscall.Flock_t{
		Start:  0,
		Len:    1,
		Type:   lockType,
		Whence: int16(os.SEEK_SET),
	}
	err = syscall.FcntlFlock(lockFile.Fd(), syscall.F_SETLK, &lock)

	return err
}
