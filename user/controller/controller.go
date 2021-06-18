package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	AddUser() interface{}
	GetAllUsers() interface{}
}

func setupResponses(code int, message error, data interface{}) map[string]interface{} {
	bodyResponses := map[string]interface{}{}

	bodyResponses["code"] = code
	bodyResponses["data"] = data

	if message != nil {
		bodyResponses["message"] = message.Error()
	} else {
		bodyResponses["message"] = "success"
	}

	return bodyResponses
}

func showResponses(code int, message error, data interface{}, ginContext *gin.Context) {
	bodyResponses := setupResponses(code, message, data)
	ginContext.JSON(code, gin.H{"responses": bodyResponses})

}
