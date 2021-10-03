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
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, "bad request")
	}
	res, errr := Service.Create(t)
	if errr != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, errr.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusCreated, "created", res)
}

func GetTransactions(c *fiber.Ctx) error {

	var res []*contract.TransactionResponse
	var err error
	custID, _ := strconv.ParseUint(c.Query("id_customer"), 10, 64)
	agentID, _ := strconv.ParseUint(c.Query("id_agen"), 10, 64)

	switch {
	case custID != 0:
		res, err = Service.GetByCustID(uint(custID))
	case agentID != 0:
		res, err = Service.GetByAgentID(uint(agentID))
		// code for agent
	default:
		res, err = Service.GetAll()
	}

	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}

func ConfirmTransactions(c *fiber.Ctx) error {
	type confirm struct {
		ID uint `json:"id_transaksi"`
	}

	idTrx := confirm{}
	err := c.BodyParser(&idTrx)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := Service.SetStatus(idTrx.ID, 1)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}

func CancelTransactions(c *fiber.Ctx) error {
	type confirm struct {
		ID uint `json:"id_transaksi"`
	}

	idTrx := confirm{}
	err := c.BodyParser(&idTrx)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := Service.SetStatus(idTrx.ID, 2)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}

func FinishTransactions(c *fiber.Ctx) error {
	type confirm struct {
		ID uint `json:"id_transaksi"`
	}

	idTrx := confirm{}
	err := c.BodyParser(&idTrx)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := Service.SetStatus(idTrx.ID, 3)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}

func RatingTransactions(c *fiber.Ctx) error {
	type rating struct {
		ID     uint `json:"id_transaksi"`
		Rating uint `json:"rating"`
	}

	trx := rating{}
	err := c.BodyParser(&trx)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := Service.SetRating(trx.ID, trx.Rating)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", res)
}

func DeleteTransactions(c *fiber.Ctx) error {
	type confirm struct {
		ID uint `json:"id_transaksi"`
	}

	idTrx := confirm{}
	err := c.BodyParser(&idTrx)

	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	err = Service.Delete(idTrx.ID)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Deleted", nil)
}
