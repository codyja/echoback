package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func formatReqHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)

	for k, _ := range r.Header {
		value := r.Header.Get(k)
		headers[k] = value
	}

	//resp := make(map[string]interface{})
	//resp["RequestHeaders"] = headers

	return headers
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]interface{})

	// format headers into map
	headers := formatReqHeaders(r)
	resp["RequestHeaders"] = headers

	// set standard headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// marshal response to json and write to ResponseWriter
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func main() {

	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}