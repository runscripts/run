package utils

import (
	"os"
	"syscall"
)

var lockFile *os.File

// Lock the file.
func Flock(path string) error {
	return fcntlFlock(syscall.F_WRLCK, path)
}

// Unlock the file.
func Funlock(path string) error {
	err := fcntlFlock(syscall.F_UNLCK)
	if err != nil {
		return err
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
			return err
		}
	}

	lock := syscall.Flock_t{
		Start:  0,
		Len:    1,
		Type:   lockType,
		Whence: int16(os.SEEK_SET),
	}
	return syscall.FcntlFlock(lockFile.Fd(), syscall.F_SETLK, &lock)
}
