package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (a AutoService) createAuto(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Params:", request.URL.Query())
	response.Header().Set("Content-Type", "application/json")
	var newAuto Auto
	newAuto.ID = request.URL.Query().Get("id")
	newAuto.Number = request.URL.Query().Get("number")
	newAuto.Model = request.URL.Query().Get("model")
	newAuto.ModelType = request.URL.Query().Get("modelType")
	// fmt.Println(request.Body)
	// params := mux.Vars(request)
	// newAuto.ID = params["id"]
	// _ = json.NewDecoder(request.Body).Decode(&newAuto)
	// fmt.Println("Auto:", newAuto)
	// fmt.Fprintf(response, "New Auto: %+v", newAuto)
	autoArray = append(autoArray, newAuto)
	json.NewEncoder(response).Encode(newAuto)
}

// // Create Single book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID just for testing...
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }
