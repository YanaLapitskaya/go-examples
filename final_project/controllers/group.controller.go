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

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Groups []models.Group `json:"groups"`
	}

	groups, err := services.GetGroups()
	if err != nil {
		panic(err)
	}

	response := Response{groups}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func AddNewGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var group models.GroupDb
	if err = json.Unmarshal(reqBody, &group); err != nil {
		panic(err)
	}

	newGroup, err := repositories.AddGroup(&group)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newGroup); err != nil {
		panic(err)
	}
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var group models.GroupDb
	if err = json.Unmarshal(reqBody, &group); err != nil {
		panic(err)
	}

	updatedGroup, err := repositories.UpdateGroup(id, &group)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updatedGroup); err != nil {
		panic(err)
	}
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := repositories.DeleteGroup(id); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
