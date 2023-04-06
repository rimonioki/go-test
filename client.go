package main

import (
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	var results []byte
	var error error
	resp, error := http.Get(url)
	if error != nil {
		return results, error
	}
	defer resp.Body.Close()
	results, error = ioutil.ReadAll(resp.Body) // response body is []byte
	return results, error
}
