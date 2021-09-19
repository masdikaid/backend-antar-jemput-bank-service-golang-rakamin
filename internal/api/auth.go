package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
	contr := contract.LoginResponse{}

	tokenya := c.Get(fiber.HeaderAuthorization)[len("Bearer "):]

	token, err := jwt.Parse(tokenya, JwtFunc)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		contr.SID = claims["sid"].(string)
	} else {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}

	err = AuthService.Logout(&contr)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success Logout", nil)
}

func JwtFunc(claim *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("SECRET")), nil
}
