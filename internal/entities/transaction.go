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
	Address     string
	Province    string
	City        string
	District    string
	Longitude   float64
	Latitude    float64
	ServicesID  uint
	Services    Services
	Type        string
	Amount      uint
	Status      uint
	Rating      float64
}
