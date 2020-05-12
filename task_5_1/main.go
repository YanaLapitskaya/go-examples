package main

import (
	"log"
	"net/http"
)

var b = []byte(`
	{
	    "host": "127.0.0.1:8080",
	    "user_agent": "curl/7.67.0",
	    "request_uri": "/anything/you?want",
	    "headers": {
	      "Accept": ["*/*"],
	      "User-Agent": ["curl/7.67.0"]
	    }
	}
`)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(b); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handleRoot)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
