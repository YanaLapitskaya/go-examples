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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var timeframe models.Timeframe
	if err = json.Unmarshal(reqBody, &timeframe); err != nil {
		panic(err)
	}

	newTimeframe, err := repositories.AddTimeframe(&timeframe)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newTimeframe); err != nil {
		panic(err)
	}
}

func DeleteTimeframe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := repositories.DeleteTimeframe(id); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
