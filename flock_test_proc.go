package main

import (
	"os"
	"time"
	"github.com/runscripts/runscripts/utils"
)

func main() {
	path := "/tmp/flock_test.lock"
	if os.Args[1] == "0" {
		if err := utils.Flock(path); err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		if err := utils.Funlock(path); err != nil {
			panic(err)
		}
	} else {
		if err := utils.Flock(path); err != nil {
			panic(err)
		}
		if err := utils.Funlock(path); err != nil {
			panic(err)
		}
	}
}
