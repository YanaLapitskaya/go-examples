package main

import (
	"encoding/json"
	"go-exercises/task_5_2/types"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var rootPath = "/"
var port = ":8081"

func handleRoot(w http.ResponseWriter, r *http.Request) {
	body := types.RequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Panic(err)
		return
	}
	value := getIntOrStringValue(body.Value)
	switch v := value.(type) {
	case int64:
		ReturnJSON(w, types.ResponseBody{Value: v * 2})
	case string:
		ReturnJSON(w, types.ResponseBody{Value: strings.ToUpper(v)})
	}
}

func getIntOrStringValue(text string) interface{} {
	var value interface{}
	number, err := strconv.ParseInt(text, 10, 64)
	if err == nil {
		value = number
	} else {
		value = text
	}
	return value
}

func main() {
	http.HandleFunc(rootPath, handleRoot)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		<-sigchan
		os.Exit(0)
	}()
}
