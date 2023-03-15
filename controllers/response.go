package controllers

import (
	"github.com/gin-gonic/gin"
)

func sendResponse(cont *gin.Context,status int, message string) {
	var response Response
	response.Status = status
	response.Message = message
	cont.IndentedJSON(response.Status,response)
}

func sendDataResponse(cont *gin.Context, status int, message string, req interface{}) {
	var response ResponseData
	response.Status = status
	response.Message = message
	response.Data = req
	cont.IndentedJSON(response.Status,response)
}