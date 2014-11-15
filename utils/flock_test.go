package utils

import "testing"
import "time"

func TestFlock(test *testing.T) {
	LogError("test logging")
	Flock("./123")
	Funlock("./123")
	time.Sleep(time.Second * 5)
}
