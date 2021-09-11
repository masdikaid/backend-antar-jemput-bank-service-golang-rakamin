package route

import (
	"backend-a-antar-jemput/internal/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// create route group api
	apiroute := app.Group("")

	// declare route
	apiroute.Get("/", api.Index)
	apiroute.Post("/transaksi/create", api.CreateTransaction)
	apiroute.Get("/transaksi/:id_customer?", api.GetTransactions)
	apiroute.Get("/transaksi/:id_agen?", api.GetTransactions)
	apiroute.Get("/transaksi", api.GetTransactions)

}
