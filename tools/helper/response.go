package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponseOkBuilder(c *fiber.Ctx, status int, msg string, data interface{}) error {
	return c.Status(status).JSON(Response{Status: status, Message: msg, Data: data})
}
