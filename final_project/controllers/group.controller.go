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

	groups := services.GetGroups()

	response := Response{groups}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AddNewGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var group models.GroupDb
	json.Unmarshal(reqBody, &group)

	newGroup := repositories.AddGroup(&group)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGroup)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var group models.GroupDb
	json.Unmarshal(reqBody, &group)

	updatedGroup := repositories.UpdateGroup(id, &group)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedGroup)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repositories.DeleteGroup(id)

	w.WriteHeader(http.StatusNoContent)
}
