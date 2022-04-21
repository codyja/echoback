package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"os"
)

func formatReqHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)

	for k, _ := range r.Header {
		value := r.Header.Get(k)
		headers[k] = value
	}

	return headers
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]interface{})

	// format headers into map
	headers := formatReqHeaders(r)
	resp["RequestHeaders"] = headers

	// if post take data in request and add to response also
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)

		resp["Body"] = string(body)
	}

	// pull environment vars and add to response also
	envs := os.Environ()
	resp["EnvironmentVariables"] = envs

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