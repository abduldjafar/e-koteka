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
			return fiberContext.Status(400).JSON(
				map[string]interface{}{
					"responses": setupResponses(400, err, nil),
				},
			)
		}
		return fiberContext.Status(200).JSON(
			map[string]interface{}{
				"responses": setupResponses(200, nil, nil),
			},
		)
	}

}

func (*fiberController) GetAllUsers() interface{} {
	return func(fiberContext *fiber.Ctx) error {
		data, err := service.FindAll()
		if err != nil {
			log.Println(err.Error())
			return fiberContext.Status(400).JSON(
				map[string]interface{}{
					"responses": setupResponses(400, err, nil),
				},
			)
		}

		return fiberContext.Status(200).JSON(
			map[string]interface{}{
				"responses": setupResponses(200, nil, data),
			},
		)

	}
}
func (*fiberController) GetUser() interface{} {
	return func(fiberContext *fiber.Ctx) error {
		id := fiberContext.Params("id")
		data, err := service.Find(id)

		if err != nil {
			log.Println(err.Error())
			return fiberContext.Status(400).JSON(
				map[string]interface{}{
					"responses": setupResponses(400, err, nil),
				},
			)
		}

		return fiberContext.Status(200).JSON(
			map[string]interface{}{
				"responses": setupResponses(200, nil, data),
			},
		)
	}
}

func FiberController() Controller {
	return &fiberController{}
}
