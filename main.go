package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestJSON struct {
	URL string
}

func DetectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request!")
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	defer r.Body.Close()
	var rj = RequestJSON{}
	err := json.Unmarshal(body, &rj)
	if err != nil {
		log.Println("error!")
	}
	reader, _ := downloadFromUrl(rj.URL)
	f, _ := faces(reader)
	log.Println(f)
}

func main() {
	http.HandleFunc("/", DetectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
