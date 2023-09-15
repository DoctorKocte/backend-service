package main

type Auto struct {
	ID string `json:"id"`
	//  номер автомобиля
	Number string `json:"number"`
	// марка автомобиля
	Model string `json:"model"`
	// модель автомобиля
	ModelType string `json:"modelType"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
