package controllers

import (
	"encoding/json"
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func AddNewTimeframe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var timeframe models.Timeframe
	json.Unmarshal(reqBody, &timeframe)

	newTimeframe := repositories.AddTimeframe(&timeframe)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTimeframe)
}

func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	repositories.DeleteTimeframe(id)

	w.WriteHeader(http.StatusNoContent)
}
