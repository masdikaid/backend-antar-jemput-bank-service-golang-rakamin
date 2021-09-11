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
	apiroute.Get("/login", api.Login)
	apiroute.Get("/logout", api.Logout)
	apiroute.Post("/transaksi/create", api.CreateTransaction)
	apiroute.Get("/transaksi/:id_customer?", api.GetAllTransactionsByCust)
	apiroute.Get("/transaksi", api.GetAllTransactions)
}
