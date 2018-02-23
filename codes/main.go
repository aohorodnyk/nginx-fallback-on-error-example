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
	codeString := strings.Trim(r.URL.String(), "/")
	codeInt, err := strconv.ParseInt(codeString, 10, 16)
	if err == nil {
		w.WriteHeader(int(codeInt))
		log["Parsed code"] = codeString
		httpMessage := "Parsed code: " + codeString
		fmt.Fprintf(w, httpMessage)
	} else {
		w.WriteHeader(500)
		log["ERROR"] = "CANNOT PARSE CODE"
		httpMessage := "ERROR: CANNOT PARSE CODE"
		fmt.Fprintf(w, httpMessage)
	}
	json, _ := json.Marshal(log)
	fmt.Println(string(json))
}
