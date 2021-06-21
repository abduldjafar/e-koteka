package main

import (
	"user/endpoints"

	"github.com/gofiber/fiber/v2"
)

var (
	Api endpoints.Endpoints = endpoints.FiberEndpoints()
)

func main() {
	api := Api.ALL().(*fiber.App)
	api.Listen(":3030")
}
