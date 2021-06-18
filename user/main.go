package main

import (
	"user/endpoints"

	"github.com/gin-gonic/gin"
)

var (
	Api endpoints.Endpoints = endpoints.NewGinEndpoints()
)

func main() {
	api := Api.ALL().(*gin.Engine)
	api.Run()
}
