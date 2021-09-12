package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
	"backend-a-antar-jemput/tools/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var Service = service.ServiceTrasaction{Repository: repository.TransactionRepositoryMysql{}}

func CreateTransaction(c *fiber.Ctx) error {
	var t *contract.Transaction
	err := c.BodyParser(&t)
	if err != nil {
		c.SendString("not found")
		return err
	}
	res, errr := Service.Create(t)
	if errr != nil {
		c.SendString("not found")
		return errr
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusCreated, "created", res)
}

func GetTransactions(c *fiber.Ctx) error {

	var res []*contract.Transaction
	var err error
	custID, _ := strconv.ParseUint(c.Query("id_customer"), 10, 64)
	agentID, _ := strconv.ParseUint(c.Query("id_agen"), 10, 64)
	switch {
	case custID != 0:
		res, err = Service.GetByCustID(uint(custID))
	case agentID != 0:
		// code for agent
	default:
		res, err = Service.GetAll()
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("not found")
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}
