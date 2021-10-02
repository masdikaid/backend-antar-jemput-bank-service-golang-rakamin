package migration

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
)

func RunMigrate() {
	// load database
	databases.Load()

	// run migration with gorm
	databases.DBCon.AutoMigrate(&entities.Login{}, &entities.Services{}, &entities.Customers{}, &entities.Agents{}, &entities.Location{}, &entities.Transaction{})
}
