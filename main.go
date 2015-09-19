package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type DetectRequest struct {
	URL string
}

func handleError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}

func DetectHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var dr DetectRequest
	err = json.Unmarshal(body, &dr)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	reader, err := downloadFromURL(dr.URL)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	f, err := faces(reader)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	http.HandleFunc("/", DetectHandler)
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", fs)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
