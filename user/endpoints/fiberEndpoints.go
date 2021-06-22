package endpoints

import (
	"os"
	"user/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	fiberRouter.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		Output: os.Stdout,
	}))

	return fiberRouter
}

func FiberEndpoints() Endpoints {
	return &fiberEndpoints{}
}
