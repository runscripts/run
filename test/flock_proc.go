package main

import (
	"os"
	"time"

	"github.com/runscripts/runscripts/utils"
)

// This program is run by flock_test.go to verify the lock effect.
func main() {
	path := "/tmp/flock_test.lock"
	if err := utils.Flock(path); err != nil {
		panic(err)
	}
	if os.Args[1] == "0" {
		time.Sleep(time.Second)
	}
	if err := utils.Funlock(path); err != nil {
		panic(err)
	}
}
