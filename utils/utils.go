package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Retrieve a file via HTTP GET.
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
		// When fetching run.conf, etc.
		return ioutil.WriteFile(path, body, 0644)
	} else {
		// When fetching scripts.
		return ioutil.WriteFile(path, body, 0777)
	}
}
