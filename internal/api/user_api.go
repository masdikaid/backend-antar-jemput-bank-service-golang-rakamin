package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"

	"github.com/gofiber/fiber/v2"
)

var ServiceAgent = service.ServiceAgent{Repository: repository.AgentRepositoryMysql{}}

func FindAgent(c *fiber.Ctx) error {
	var body *contract.Customer
	err := c.BodyParser(&body)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "bad request")
	}

	list, err2 := ServiceAgent.GetListAgent(body.District, body.City)
	if err2 != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "bad request")
	}

	body.ListAgent = list
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "found", body)
}
