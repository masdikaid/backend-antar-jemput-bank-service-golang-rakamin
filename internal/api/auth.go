package api

import (
	"github.com/gofiber/fiber/v2"
)

// example route for login
func Login(c *fiber.Ctx) error {
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  fiber.StatusOK,
		Message: "berhasil login",
	}
	return c.JSON(response)
}

func Logout(c *fiber.Ctx) error {
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  fiber.StatusOK,
		Message: "Berhasil logout",
	}
	return c.JSON(response)
}
