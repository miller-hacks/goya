package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func downloadFromURL(url string) (*bytes.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while downloading %s: %s", url, err))
	}
	defer response.Body.Close()

	imageBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error casting Closer to bytes")
	}
	return bytes.NewReader(imageBytes), nil
}
