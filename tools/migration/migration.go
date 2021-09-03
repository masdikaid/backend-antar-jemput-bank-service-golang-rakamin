package migration

import (
	"bri-antar-jemput/internal/databases"
	"bri-antar-jemput/internal/models"
)

func RunMigrate() {
	// load database
	databases.Load()

	// run migration with gorm
	databases.DBCon.AutoMigrate(&models.Login{})
}
