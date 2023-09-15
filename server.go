package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var autoArray []Auto

func main() {
	var router = mux.NewRouter()

	var autoService AutoServiceInterface = AutoService{}
	// mock
	autoArray = append(autoArray, Auto{ID: "1", Number: "4366IB1", Model: "Volkswagen", ModelType: "Arteon"})

	router.HandleFunc("/allAuto", autoService.getAutoArray).Methods("GET")
	router.HandleFunc("/auto/{id}", autoService.getAuto).Methods("GET")
	router.HandleFunc("/createAuto", autoService.createAuto).Methods("POST")
	router.HandleFunc("/updateAuto/{id}", autoService.updateAuto).Methods("PUT")
	router.HandleFunc("/deleteAuto/{id}", autoService.deleteAuto).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
