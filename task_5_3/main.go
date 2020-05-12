package main

import (
	"fmt"
	"log"
	"net/http"
)

type PostRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var port string = ":8081"
var staticDir string = "./static"
var rootPath string = "/"

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != rootPath || (r.Method != http.MethodGet && r.Method != http.MethodPost) {
		http.NotFound(w, r)
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, staticDir)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		cookie := http.Cookie{Name: "token", Value: fmt.Sprintf("%v:%v", r.Form["name"][0], r.Form["address"][0])}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, rootPath, http.StatusFound)
	}
}

func main() {
	http.HandleFunc(rootPath, handleRoot)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}
}
