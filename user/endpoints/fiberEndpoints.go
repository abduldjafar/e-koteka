package endpoints

import (
	"user/controller"

	"github.com/gofiber/fiber/v2"
)

type fiberEndpoints struct{}

var (
	fiberRouter = fiber.New(
		fiber.Config{
			Concurrency: 1000000,
		},
	)
	fiberUserController controller.Controller = controller.FiberController()
)

func (*fiberEndpoints) ALL() interface{} {
	fiberRouter.Post("/v1/users", fiberUserController.AddUser().(fiber.Handler))
	fiberRouter.Get("/v1/users/all", fiberUserController.GetAllUsers().(fiber.Handler))
	fiberRouter.Get("/v1/users/:id", fiberUserController.GetUser().(fiber.Handler))

	return fiberRouter
}

func FiberEndpoints() Endpoints {
	return &fiberEndpoints{}
}
