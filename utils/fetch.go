package utils

import (
	"io/ioutil"
	"net/http"
)

// Http Get to fetch file.
func Fetch(url string, path string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return Errorf("%s: %s", response.Status, url)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, body, 0777)
}
