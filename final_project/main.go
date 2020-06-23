package main

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type PostRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

var appPort string = ":8080"
var staticDir string = "./static"
var rootPath string = "/"

func main() {

	configs.InitDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.AddNewTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	if err := http.ListenAndServe(appPort, router); err != nil {
		log.Panic(err)
	}
}
