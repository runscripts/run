package flock

import (
	"os/exec"
	"testing"
)

func TestFlock(t *testing.T) {
	path := "/tmp/flock_test.lock"
	err := Flock(path)
	if err != nil {
		t.Fatal(err)
	}
	err = Funlock(path)
	if err != nil {
		t.Fatal(err)
	}
	cmd0 := exec.Command("go", "run", "../test/flock_proc.go", "0")
	err = cmd0.Start()
	if err != nil {
		t.Fatal(err)
	}
	cmd1 := exec.Command("go", "run", "../test/flock_proc.go", "1")
	err = cmd1.Start()
	if err != nil {
		t.Fatal(err)
	}
	err = cmd1.Wait()
	if err == nil {
		t.Errorf("%s shouldn't be locked twice", path)
	}
	err = cmd0.Wait()
	if err != nil {
		t.Error(err)
	}
}
