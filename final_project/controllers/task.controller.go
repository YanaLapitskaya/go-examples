package controllers

import (
	"encoding/json"
	"fmt"
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"go-exercises/final_project/services"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Tasks []models.Task `json:"tasks"`
	}

	tasks := services.GetTasks()

	response := Response{tasks}
	fmt.Println(response)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AddNewTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.TaskDb
	json.Unmarshal(reqBody, &task)

	newTask := repositories.AddTask(&task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	json.Unmarshal(reqBody, &task)

	updatedTask := repositories.UpdateTask(id, &task)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repositories.DeleteTask(id)

	w.WriteHeader(http.StatusNoContent)
}
