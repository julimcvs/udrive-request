package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"udrive-request/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var request model.Request

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	id, err := model.Insert(request)

	var msg string
	var status int

	if err != nil {
		msg = fmt.Sprintf("Error creating ride: %v", err)
		status = 400
	} else {
		msg = fmt.Sprintf("Ride successfully created! ID: %v", id)
		status = 201
	}

	responseBody := model.ResponseBody{
		Status:  &status,
		Message: &msg,
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&responseBody)

	if err != nil {
		return
	}
}
