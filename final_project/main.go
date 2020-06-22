package main

import (
	"encoding/json"
	"fmt"
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"io/ioutil"
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

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Tasks []models.Task `json:"tasks"`
	}
	rows := repositories.GetAllTasks()
	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	response := Response{tasks}
	fmt.Println(response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func addNewTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)

	newTask := repositories.AddTask(&task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)

	updatedTask := repositories.UpdateTask(id, &task)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repositories.DeleteTask(id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {

	configs.InitDB()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", addNewTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	if err := http.ListenAndServe(appPort, router); err != nil {
		log.Panic(err)
	}
}
