package route

import (
	"backend-a-antar-jemput/internal/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// create route group api
	apiroute := app.Group("api")

	// declare route
	apiroute.Get("/", api.Index)
	apiroute.Get("/transaksi", api.GetAllTransactions)

}
