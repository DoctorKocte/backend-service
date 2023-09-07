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
	autoArray = append(autoArray, Auto{ID: "1", Number: "4366IB1", Model: "Volkswagen", ModelType: "Arteon"})

	router.HandleFunc("/api/allAuto", autoService.getAutoArray).Methods("GET")
	router.HandleFunc("/api/auto/{id}", autoService.getAuto).Methods("GET")
	//router.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
