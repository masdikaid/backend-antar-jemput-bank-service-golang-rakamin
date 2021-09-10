package migration

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/models"
)

func RunMigrate() {
	// load database
	databases.Load()

	// run migration with gorm
	databases.DBCon.AutoMigrate(&models.Login{},&models.Users{},&models.Customers{},&models.Agents{},&models.Location{},&models.Transaction{})
}
