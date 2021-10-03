package main

import (
	// project import
	_ "backend-a-antar-jemput/config"
	//"backend-a-antar-jemput/internal/models"
	"backend-a-antar-jemput/internal/route"
	"backend-a-antar-jemput/tools/migration"
	"backend-a-antar-jemput/tools/seeds"

	//"fmt"

	// default import
	"log"
	"os"

	// package import
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// check if migration call in arg
	// its not uptodate

	args := os.Args[1:]
	// when main call with argument
	if len(args) > 0 {
		switch args[0] {
		case "db:migrate":
			migration.RunMigrate()
		case "db:seeds":
			seeds.RunSeeds()
		default:
			log.Fatalln("argument not found !")
		}
	} else {
		// a := models.Location{}
		// fmt.Println(a.GetLocationByLogin())

		// setup fiber app
		app := SetupFiber()

		// listening
		log.Fatal(app.Listen(":" + os.Getenv("PORT")))
	}

}

func SetupFiber() *fiber.App {
	// init http with fiber
	app := fiber.New(fiber.Config{})
	// setup log
	app.Use(logger.New(logger.ConfigDefault))
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://127.0.0.1:3000, http://localhost:3000",
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
	// 	AllowCredentials: true,
	// }))
	// setup route
	route.SetupRoutes(app)
	return app
}
