package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	var log map[string]string
	log = make(map[string]string)
	log["Source URL"] = r.URL.String()
	path := strings.Trim(r.URL.String(), "/")
	pathInt64, _ := strconv.ParseInt(path, 10, 32)
	code := int(pathInt64)
	statusText := http.StatusText(code)
	if statusText != "" {
		log["Parsed Code"] = path
		log["Status Text"] = statusText
	} else {
		code = 500
		log["ERROR"] = "CANNOT PARSE CODE"
	}
	json, _ := json.Marshal(log)
	jsonMessage := string(json)

	w.WriteHeader(code)
	fmt.Fprintf(w, jsonMessage)

	fmt.Println(jsonMessage)
}
