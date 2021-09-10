package api

import (
	//"backend-a-antar-jemput/internal/models"

	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetAllTransactions(c *fiber.Ctx) error {
	// response := struct {
	// 	Status  int    `json:"status"`
	// 	Message string `json:"message"`
	// }{
	// 	Status:  fiber.StatusOK,
	// 	Message: "Success Access Transaksi Ok",
	// }
	// return c.JSON(response)
	var trx []contract.Transaction
	var err error
	trx = service.GetAllTransactions()
	//var err error
	//x := models.GetAll()
	//trx,err = x,nil
	if err != nil {
		panic(err)
	}
	return c.JSON(trx)
}

func GetAllTransactionsCust(c *fiber.Ctx) error  {
	var trx []contract.TransactionCust
	var err error
	cust := c.Params("id")
	trx = service.GetAllTransactionsCust(cust)
	if err != nil {
		panic(err)
	}
	return c.JSON(trx)
}
