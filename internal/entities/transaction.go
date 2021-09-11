package entities

import (
	//"backend-a-antar-jemput/internal/databases"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomersID uint
	Customers   Customers
	AgentsID    uint
	Agents      Agents
	LocationID  uint
	Location    Location
	Tipe        string
	Amount      int
	Status      int
	Rating      float64
}
