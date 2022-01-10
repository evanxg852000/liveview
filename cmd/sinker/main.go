package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
		Status  string
	}{
		Message: "Welcome to sinker",
		Status:  "running",
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func IngestDataHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := make(map[string](interface{}))
	err = json.Unmarshal(payload, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Println(data)
}

func main() {
	http.HandleFunc("/", StatusHandler)
	http.HandleFunc("/ingest", IngestDataHandler)

	fmt.Println("Sinker started on port 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
