package utils

import (
//	"io/ioutil"
//	"net/http"
	"os"
)

func Fetch(url string, dir string, name string) error {
	return nil
}

func MakeDir(path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		LogError("cannot mkdir %s\n", path)
		panic(err)
	}
}
