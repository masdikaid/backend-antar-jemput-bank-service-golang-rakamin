package route

import (
	"backend-a-antar-jemput/internal/api"
	"backend-a-antar-jemput/internal/midleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// create route group api
	app.Post("/login", api.Login)
	apiroute := app.Group("", midleware.AuthMiddleware())

	// declare route
	apiroute.Get("/", api.Index)
	apiroute.Get("/logout", api.Logout)
	apiroute.Post("/cariagen", api.FindAgent)
	apiroute.Get("/profil", api.GetProfile)
	apiroute.Get("/transaksi", api.GetTransactions)
	apiroute.Post("/transaksi/create", api.CreateTransaction)
	apiroute.Post("/transaksi/dikonfirmasi", api.ConfirmTransactions)
	apiroute.Post("/transaksi/dibatalkan", api.CancelTransactions)
	apiroute.Post("/transaksi/selesai", api.FinishTransactions)
	apiroute.Delete("/transaksi/delete/:id", api.DeleteTransactions)
}
