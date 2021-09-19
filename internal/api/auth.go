package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"

	"github.com/gofiber/fiber/v2"
)

var AuthService = service.AuthService{AuthRepository: repository.AuthRepositoryMysql{}, SessionRepository: repository.SessionRepositoryMysql{}}

// example route for login
func Login(c *fiber.Ctx) error {
	var contr contract.LoginRequest
	err := c.BodyParser(&contr)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "Bad Request")
	}
	var token string
	token, err = AuthService.Login(contr)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusAccepted, "success", token)
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
