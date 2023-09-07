package main

import (
	"net/http"
)

type Auto struct {
	ID string `json:"id"`
	//  номер автомобиля
	Number string `json:"number"`
	// марка автомобиля
	Model string `json:"model"`
	// модель автомобиля
	ModelType string `json:"modelType"`
}

type AutoServiceInterface interface {
	getAutoArray(respose http.ResponseWriter, _ *http.Request)
	getAuto(response http.ResponseWriter, request *http.Request)
}
