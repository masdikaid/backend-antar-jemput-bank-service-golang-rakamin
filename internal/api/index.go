package api

import "github.com/gofiber/fiber/v2"

// example route
func Index(c *fiber.Ctx) error {
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  fiber.StatusOK,
		Message: "Success Access OK",
	}
	return c.JSON(response)
}
