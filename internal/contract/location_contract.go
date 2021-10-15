package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/tools/helper"
)

type Location struct {
	ID       uint
	Address  string
	Province string
	City     string
	District string
}

func (l *Location) ToEntity(loginID uint) *entities.Location {
	ent := entities.Location{}
	helper.ConvertStruct(l, ent)
	ent.Province = l.Province
	ent.City = l.City
	ent.District = l.District
	ent.Address = l.Address
	ent.LoginID = loginID
	return &ent
}

func (l *Location) FromEntity(source *entities.Location) {
	helper.ConvertStruct(source, l)
}

type Province struct {
	Id   string
	Name string
}

type City struct {
	Id          string
	Province_id string
	Name        string
}

type District struct {
	Id         string
	Regency_id string
	Name       string
}
