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
	apiroute.Post("/cariagen", api.FindAgent)
	apiroute.Get("/transaksi", api.GetTransactions)
	apiroute.Post("/transaksi/create", api.CreateTransaction)
	apiroute.Post("/transaksi/dikonfirmasi", api.ConfirmTransactions)
	apiroute.Post("/transaksi/dibatalkan", api.CancelTransactions)
	apiroute.Post("/transaksi/selesai", api.FinishTransactions)
	apiroute.Delete("/transaksi/delete/:id", api.DeleteTransactions)
}
