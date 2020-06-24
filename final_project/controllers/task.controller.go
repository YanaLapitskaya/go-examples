package controllers

import (
	"encoding/json"
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

	tasks, err := services.GetTasks()
	if err != nil {
		panic(err)
	}

	response := Response{tasks}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func AddNewTask(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var task models.TaskDb
	if err = json.Unmarshal(reqBody, &task); err != nil {
		panic(err)
	}

	newTask, err := repositories.AddTask(&task)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newTask); err != nil {
		panic(err)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task models.Task
	if err := json.Unmarshal(reqBody, &task); err != nil {
		panic(err)
	}

	updatedTask, err := repositories.UpdateTask(id, &task)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updatedTask); err != nil {
		panic(err)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := repositories.DeleteTask(id); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
