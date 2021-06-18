package endpoints

import (
	"user/controller"

	"github.com/gin-gonic/gin"
)

type ginEndpoints struct{}

var (
	ginRouter                               = gin.Default()
	ginUserController controller.Controller = controller.NewUserController()
)

func (*ginEndpoints) ALL() interface{} {

	ginRouter.POST("/v1/users", ginUserController.AddUser().(func(*gin.Context)))
	ginRouter.GET("/v1/users/all", ginUserController.GetAllUsers().(func(*gin.Context)))

	return ginRouter

}
func NewGinEndpoints() Endpoints {
	return &ginEndpoints{}
}
