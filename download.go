package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func downloadFromUrl(url string) (*bytes.Reader, error) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	output, err := os.Create(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while creating", fileName, "-", err))
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while downloading", url, "-", err))
	}
	defer response.Body.Close()

	// n, err := io.Copy(output, response.Body)
	// if err != nil {
	// 	fmt.Println("Error while downloading", url, "-", err)
	// 	return
	// }

	// fmt.Println(n, "bytes downloaded.")

	imageBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error casting Closer to bytes")
	}
	imageReader := bytes.NewReader(imageBytes)

	return imageReader, nil
}
