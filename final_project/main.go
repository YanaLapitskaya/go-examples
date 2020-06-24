package main

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/controllers"
	"go-exercises/final_project/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var appPort string = ":8080"
var staticDir string = "./static"
var rootPath string = "/"

func main() {

	configs.InitDB()
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.LoggingHandler)
	router.Use(middlewares.ErrorHandler)

	router.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.AddNewTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	router.HandleFunc("/groups", controllers.GetAllGroups).Methods("GET")
	router.HandleFunc("/groups", controllers.AddNewGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", controllers.UpdateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", controllers.DeleteGroup).Methods("DELETE")

	router.HandleFunc("/timeframes", controllers.AddNewTimeframe).Methods("POST")
	router.HandleFunc("/timeframes/{id}", controllers.DeleteTimeframe).Methods("DELETE")

	if err := http.ListenAndServe(appPort, router); err != nil {
		log.Panic(err)
	}
}
