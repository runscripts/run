package utils

import (
	"io/ioutil"
	"net/http"
)

// Http Get to fetch file.
func Fetch(url string, path string) error {
	response, err := http.Get(url)
	if err != nil {
		LogError("Failed to GET %s\n", url)
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return ioutil.WriteFile(path, body, 0777)
}
