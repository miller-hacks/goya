package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

type DetectRequest struct {
	URL string `json:"url"`
}

type DetectResponse struct {
	Faces []*Face `json:"faces"`
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
	resp := &DetectResponse{
		Faces: f,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	http.HandleFunc("/", DetectHandler)
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	log.Println("Server started")
	var ip_port = flag.String("b", ":8000", "address:port")
	flag.Parse()
	log.Fatal(http.ListenAndServe(*ip_port, nil))
}
