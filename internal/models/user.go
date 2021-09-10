package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name        string
	PhoneNumber string
}

type Customers struct {
	User    Users `gorm:"embedded;embeddedPrefix:customer_"`
	LoginID uint
	Login   Login
}

type Agents struct {
	User       Users `gorm:"embedded;embeddedPrefix:agent_"`
	OutletName string
	LoginID    uint
	Login      Login
	LocationID uint
	Location   Location
}

type UserInterface interface {
	Create() (*UserInterface, error)
	GetAll() ([]*UserInterface, error)
	GetByID() (UserInterface, error)
}

// func (c *Customers) GetByID() (*Customers, error) {
// 	var cust Customers
// 	databases.DBCon.First(&cust, "id = ?", 1).Scan(&cust)
// 	if cust.User.ID == 0 {
// 		return nil, gorm.ErrEmptySlice
// 	}
// 	return &cust, nil
// }