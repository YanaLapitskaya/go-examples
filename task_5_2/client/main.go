package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go-exercises/task_5_2/types"
	"log"
	"net/http"
	"os"
)

var exitText = "exit"
var serverUrl = "http://localhost:8081"
var jsonContentType = "application/json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == exitText {
			break
		}
		sendRequest(text)
	}

}

func sendRequest(text string) {
	reqBody, jsonErr := json.Marshal(map[string]interface{}{
		"value": text,
	})
	if jsonErr != nil {
		log.Panic(jsonErr)
	}
	resp, resErr := http.Post(serverUrl, jsonContentType, bytes.NewBuffer(reqBody))
	if resErr != nil {
		log.Panic(resErr)
	}
	body := types.ResponseBody{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Panic(err)
		return
	}
	fmt.Println(body.Value)
}
