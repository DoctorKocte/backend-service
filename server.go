package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var autoArray []Auto

type AutoService struct{}

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

func main() {
	var router = mux.NewRouter()

	var autoService AutoServiceInterface = AutoService{}
	autoArray = append(autoArray, Auto{ID: "1", Number: "4366IB1", Model: "Volkswagen", ModelType: "Arteon"})

	//http.ListenAndServe()
	router.HandleFunc("/api/allAuto", autoService.getAutoArray).Methods("GET")
	router.HandleFunc("/api/auto/{id}", autoService.getAuto).Methods("GET")
	//router.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
