package controller

import (
	"log"
	"user/entity"

	"github.com/gofiber/fiber/v2"
)

type fiberController struct{}

func (*fiberController) AddUser() interface{} {
	return func(fiberContext *fiber.Ctx) error {
		var user entity.CustomerUser

		if err := fiberContext.BodyParser(&user); err != nil {
			log.Println(err.Error())
			return fiberContext.Status(400).JSON(
				map[string]interface{}{
					"responses": setupResponses(400, err, nil),
				},
			)
		}

		if err := service.Create(user); err != nil {
			log.Println(err.Error())
			return fiberShowresponses(400, err, nil, fiberContext)
		}
		return fiberShowresponses(200, nil, nil, fiberContext)
	}

}

func (*fiberController) GetAllUsers() interface{} {
	return func(fiberContext *fiber.Ctx) error {
		data, err := service.FindAll()
		if err != nil {
			log.Println(err.Error())
			return fiberShowresponses(400, err, nil, fiberContext)
		}

		return fiberShowresponses(400, nil, data, fiberContext)
	}
}
func (*fiberController) GetUser() interface{} {
	return func(fiberContext *fiber.Ctx) error {
		id := fiberContext.Params("id")
		data, err := service.Find(id)

		if err != nil {
			log.Println(err.Error())
			return fiberShowresponses(400, err, nil, fiberContext)
		}

		return fiberShowresponses(200, nil, data, fiberContext)
	}
}

func FiberController() Controller {
	return &fiberController{}
}
