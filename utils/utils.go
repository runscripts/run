package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
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
	if strings.HasPrefix(url, MASTER_URL) {
		// Downloaded run.conf, etc.
		return ioutil.WriteFile(path, body, 0644)
	} else {
		// Downloaded scripts.
		return ioutil.WriteFile(path, body, 0777)
	}
}
