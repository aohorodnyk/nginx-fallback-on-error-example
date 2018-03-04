package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", httpHandler)
	fmt.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	var logMessage map[string]string
	logMessage = make(map[string]string)
	logMessage["Source URL"] = r.URL.String()
	path := strings.Trim(r.URL.String(), "/")
	pathInt64, _ := strconv.ParseInt(path, 10, 32)
	code := int(pathInt64)
	statusText := http.StatusText(code)
	if statusText != "" {
		logMessage["Parsed Code"] = path
		logMessage["Status Text"] = statusText
	} else {
		code = 500
		logMessage["ERROR"] = "CANNOT PARSE CODE"
	}
	json, _ := json.Marshal(logMessage)
	jsonMessage := string(json)

	w.WriteHeader(code)
	fmt.Fprintf(w, jsonMessage)

	fmt.Println(jsonMessage)
}
