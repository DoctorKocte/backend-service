package main

import (
	"net/http"
)

type AutoServiceInterface interface {
	getAutoArray(respose http.ResponseWriter, _ *http.Request)
	getAuto(response http.ResponseWriter, request *http.Request)
	createAuto(response http.ResponseWriter, request *http.Request)
	updateAuto(request http.ResponseWriter, response *http.Request)
	deleteAuto(response http.ResponseWriter, request *http.Request)
}

type AutoService struct{}
