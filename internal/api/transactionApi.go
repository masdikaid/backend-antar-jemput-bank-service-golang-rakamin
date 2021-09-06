package api

import (
	//"backend-a-antar-jemput/internal/models"

	"github.com/gofiber/fiber/v2"
)


func GetAll(c *fiber.Ctx) error {
	response := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  fiber.StatusOK,
		Message: "Success Access Transaksi Ok",
	}
	return c.JSON(response)

	// var trx []models.Transaction
	// var err error
	// x := models.GetAll()
	// trx,err = x,nil
	// if err !=nil{
	// 	panic(err)
	// }
	// return c.JSON(trx)

}