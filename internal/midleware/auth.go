package midleware

import (
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(os.Getenv("SECRET")),
		SuccessHandler: ValidateSession,
		ErrorHandler:   jwtError,
	})
}

func ValidateSession(c *fiber.Ctx) error {
	services := service.AuthService{AuthRepository: repository.AuthRepositoryMysql{}, SessionRepository: repository.SessionRepositoryMysql{}}
	strToken := c.Get(fiber.HeaderAuthorization)[len("Bearer "):]
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return jwtError(c, err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessid := claims["sid"]
		err = services.ValidateSession(sessid.(string))
		if err != nil {
			return jwtError(c, err)
		}
		return c.Next()
	}
	return jwtError(c, errors.New("invalid token"))
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return helper.JsonResponseFailBuilder(c, fiber.StatusForbidden, err.Error())
	}
	return helper.JsonResponseFailBuilder(c, fiber.StatusForbidden, err.Error())
}
