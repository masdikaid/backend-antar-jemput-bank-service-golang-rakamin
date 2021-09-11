package api

import (
	//"backend-a-antar-jemput/internal/models"

	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/internal/service"
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
	return c.JSON(res)
}

func GetAllTransactions(c *fiber.Ctx) error {
	res, err := Service.GetAll()
	if err != nil {
		c.SendString("not found")
		return err
	}
	return c.JSON(res)
}

func GetAllTransactionsByCust(c *fiber.Ctx) error {
	id := c.Query("id_customer")
	parsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.SendString("id not int")
		return err
	}
	res, err := Service.GetByCustID(uint(parsed))
	if err != nil {
		c.SendString("not found")
		return err
	}
	return c.JSON(res)
}

// func GetAllTransactions(c *fiber.Ctx) error {
// 	// response := struct {
// 	// 	Status  int    `json:"status"`
// 	// 	Message string `json:"message"`
// 	// }{
// 	// 	Status:  fiber.StatusOK,
// 	// 	Message: "Success Access Transaksi Ok",
// 	// }
// 	// return c.JSON(response)
// 	var trx []contract.Transaction
// 	var err error
// 	trx = service.GetAllTransactions()
// 	//var err error
// 	//x := models.GetAll()
// 	//trx,err = x,nil
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.JSON(trx)
// }

// func GetAllTransactionsCust(c *fiber.Ctx) error  {
// 	var trx []contract.TransactionCust
// 	var err error
// 	cust := c.Params("id")
// 	trx = service.GetAllTransactionsCust(cust)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.JSON(trx)
// }
