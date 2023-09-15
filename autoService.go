package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (a AutoService) getAutoArray(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(autoArray)
}

func (a AutoService) getAuto(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range autoArray {
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
}

func (a AutoService) updateAuto(request http.ResponseWriter, response *http.Request) {
	request.Header().Set("Content-Type", "application/json")
	params := mux.Vars(response)

	for index, item := range autoArray {
		if item.ID == params["id"] {
			autoArray = append(autoArray[:index], autoArray[index+1:]...)

			var updateAuto Auto
			_ = json.NewDecoder(response.Body).Decode(&updateAuto)
			updateAuto.ID = params["id"]
			autoArray = append(autoArray, updateAuto)
			json.NewEncoder(request).Encode(autoArray)
			return
		}
	}
}

func (a AutoService) createAuto(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var newAuto Auto
	_ = json.NewDecoder(request.Body).Decode(&newAuto)
	if len(autoArray) == 0 {
		newAuto.ID = "1"
		autoArray = append(autoArray, newAuto)
		json.NewEncoder(response).Encode(newAuto)
	} else {
		var lastID = autoArray[len(autoArray)-1].ID
		number, _ := strconv.Atoi(lastID)
		newAuto.ID = fmt.Sprintf("%v", number+1)
		autoArray = append(autoArray, newAuto)
		json.NewEncoder(response).Encode(newAuto)
	}
}

func (a AutoService) deleteAuto(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range autoArray {
		if item.ID == params["id"] {
			autoArray = append(autoArray[index+1:], autoArray[:index]...)
			json.NewEncoder(response).Encode(item)
			return
		}
	}
}
