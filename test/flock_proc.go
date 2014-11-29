package main

import (
	"os"
	"time"

	"github.com/runscripts/run/utils"
)

// This program is run by flock_test.go to verify the lock effect.
func main() {
	path := "/tmp/flock_test.lock"
	if os.Args[1] == "1" {
		time.Sleep(time.Millisecond * 100)
	}
	if err := utils.Flock(path); err != nil {
		panic(err)
	}
	if os.Args[1] == "0" {
		time.Sleep(time.Millisecond * 200)
		if err := utils.Funlock(path); err != nil {
			panic(err)
		}
	}
}
