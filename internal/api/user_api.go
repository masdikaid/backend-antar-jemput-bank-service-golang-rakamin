package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var ServiceAgent = service.ServiceAgent{Repository: repository.NewAgentRepo(databases.DBCon)}
var ServiceCustomer = service.ServiceCustomer{Repository: repository.CustomerRepositoryMysql{}}

func FindAgent(c *fiber.Ctx) error {
	var body *contract.Customer
	err := c.BodyParser(&body)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "bad request")
	}

	list, errr := ServiceAgent.GetListAgent(body.Service, body.District, body.Amount)
	if errr != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "bad request")
	}

	body.ListAgent = list
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "found", body)
}

func GetProfile(c *fiber.Ctx) error {
	var agent *contract.DetailAGent
	var customer *contract.Customer
	var err error
	custID, _ := strconv.ParseUint(c.Query("id_customer"), 10, 64)
	agentID, _ := strconv.ParseUint(c.Query("id_agen"), 10, 64)

	switch {
	case custID != 0:
		customer, err = ServiceCustomer.GetCustomer(uint(custID))
		if err != nil {
			return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
		}
		return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", customer)

	case agentID != 0:
		agent, err = ServiceAgent.GetAgent(uint(agentID))
		if err != nil {
			return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
		}

		list, errr := ServiceAgent.GetAgentService(uint(agentID))
		if errr != nil {
			return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, errr.Error())
		}

		agent.Service = list
		return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", agent)

	default:
		return helper.JsonResponseFailBuilder(c, fiber.StatusNotFound, err.Error())
	}
}

// func CreateProfileAgent(c *fiber.Ctx) error {
// 	var agent *contract.DetailAGent
// 	err := c.BodyParser(&agent)
// 	if err != nil {
// 		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
// 	}

// 	_, errr := ServiceAgent.Create(agent)
// 	if errr != nil {
// 		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, errr.Error())
// 	}

// 	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "created", agent)
// }
