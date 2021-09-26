package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name        string
	PhoneNumber string
}

type Customers struct {
	Users
	LoginID uint
	Login   Login
}

type Agents struct {
	Users
	OutletName  string
	IsAvailable bool
	MaxTrx      int
	Rating      float64
	LoginID     uint
	Login       Login
	LocationID  uint
	Location    Location
	Services    []*Services `gorm:"many2many:agent_services;"`
}

type Services struct {
	gorm.Model
	ServiceName     string
	TransactionName string
	Agents          []*Agents `gorm:"many2many:agent_services;"`
}
