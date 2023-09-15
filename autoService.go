package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	existError string = "Element with such id doesn't exist"
)

func (a AutoService) getAutoArray(response http.ResponseWriter, request *http.Request) {
	log.Printf("Request from %s for %s", request.RemoteAddr, request.URL.Path)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(autoArray)
}

func (a AutoService) getAuto(response http.ResponseWriter, request *http.Request) {
	log.Printf("Request from %s for %s", request.RemoteAddr, request.URL.Path)
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	isElementExist := false

	for _, item := range autoArray {
		if item.ID == params["id"] {
			isElementExist = true
			json.NewEncoder(response).Encode(item)
			return
		}
	}

	if !isElementExist {
		log.Println("Element with such id doesn't exist")
		json.NewEncoder(response).Encode(ErrorMessage{Message: existError})
	}
}

func (a AutoService) updateAuto(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	isElementExist := false

	for index, item := range autoArray {
		if item.ID == params["id"] {
			isElementExist = true
			autoArray = append(autoArray[:index], autoArray[index+1:]...)

			var updateAuto Auto
			_ = json.NewDecoder(request.Body).Decode(&updateAuto)
			updateAuto.ID = params["id"]
			autoArray = append(autoArray, updateAuto)
			json.NewEncoder(response).Encode(autoArray)
			return
		}
	}

	if !isElementExist {
		log.Println("Element with such id doesn't exist")
		json.NewEncoder(response).Encode(ErrorMessage{Message: existError})
	}
}

func (a AutoService) createAuto(response http.ResponseWriter, request *http.Request) {
	log.Printf("Request from %s for %s", request.RemoteAddr, request.URL.Path)

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
	log.Printf("Request from %s for %s", request.RemoteAddr, request.URL.Path)
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	isElementExist := false

	for index, item := range autoArray {
		if item.ID == params["id"] {
			isElementExist = true
			autoArray = append(autoArray[index+1:], autoArray[:index]...)
			json.NewEncoder(response).Encode(item)
			return
		}
	}

	if !isElementExist {
		log.Println("Element with such id doesn't exist")
		json.NewEncoder(response).Encode(ErrorMessage{Message: existError})
	}
}
