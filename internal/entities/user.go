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
	OutletName string
	LoginID    uint
	Login      Login
	LocationID uint
	Location   Location
}
