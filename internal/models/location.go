package models

import "gorm.io/gorm"

type LocationInterface interface {
	Create() (*LocationInterface, error)
	GetLocationByLogin(login Login) ([]*LocationInterface, error)
}

type Location struct {
	gorm.Model
	LoginID uint
	Login Login
	Address    string
	Province   string
	City       string
	District   string
	Latitude   float64
	Longitude  float64
}

func (l *Location) Create() (*Location, error) {
	return l, nil
}

func (l *Location) GetLocationByLogin() ([]Location, error) {
	return nil, nil
}