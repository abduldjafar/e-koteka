package controller

import (
	"log"
	"net/http"
	"user/entity"
	"user/services"

	"github.com/gin-gonic/gin"
)

type userController struct{}

var (
	service services.Services = services.NewServices()
)

func (*userController) AddUser() interface{} {
	return func(ginContext *gin.Context) {
		var user entity.CustomerUser

		if err := ginContext.ShouldBindJSON(&user); err != nil {
			log.Println(err.Error())
			showResponses(http.StatusBadRequest, err, nil, ginContext)
			return
		}

		if err := service.Create(user); err != nil {
			log.Println(err.Error())
			showResponses(http.StatusBadRequest, err, nil, ginContext)
			return
		}
		showResponses(http.StatusOK, nil, nil, ginContext)
	}

}

func (*userController) GetAllUsers() interface{} {
	return func(ginContext *gin.Context) {
		data, err := service.FindAll()
		if err != nil {
			log.Println(err.Error())
			showResponses(http.StatusBadRequest, err, nil, ginContext)
			return
		}

		showResponses(http.StatusOK, nil, data, ginContext)
	}
}
func (*userController) GetUser() interface{} {
	return func(ginContext *gin.Context) {
		id := ginContext.Param("id")
		data, err := service.Find(id)

		if err != nil {
			log.Println(err.Error())
			showResponses(http.StatusBadRequest, err, nil, ginContext)
			return
		}

		showResponses(http.StatusOK, nil, data, ginContext)

	}
}

func (*userController) UpdateUser() interface{} {
	return nil
}
func NewUserController() Controller {
	return &userController{}
}
