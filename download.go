package main

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	// n, err := io.Copy(output, response.Body)
	// if err != nil {
	// 	fmt.Println("Error while downloading", url, "-", err)
	// 	return
	// }

	fmt.Println(n, "bytes downloaded.")

	return response.Body
}

func main() {
	downloadFromUrl("https://pp.vk.me/c622920/v622920070/46cb0/otOfIiFiXik.jpg")
}
