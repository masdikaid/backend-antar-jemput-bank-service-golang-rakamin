package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFail struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func JsonResponseOkBuilder(c *fiber.Ctx, status int, msg string, data interface{}) error {
	return c.Status(status).JSON(Response{Status: status, Message: msg, Data: data})
}

func JsonResponseFailBuilder(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(ResponseFail{Status: status, Message: msg})
}
