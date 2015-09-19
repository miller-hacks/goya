package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func downloadFromUrl(url string) (*bytes.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while downloading", url, "-", err))
	}
	defer response.Body.Close()

	imageBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error casting Closer to bytes")
	}
	imageReader := bytes.NewReader(imageBytes)

	return imageReader, nil
}

func main() {
	downloadFromUrl("https://pp.vk.me/c622920/v622920070/46cb0/otOfIiFiXik.jpg")
}
