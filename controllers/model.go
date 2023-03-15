package controllers

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Rating   int    `json:"rating"`
	Director string `json:"director"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}