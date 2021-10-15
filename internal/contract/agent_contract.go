package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/tools/helper"
)

type ListAgent struct {
	ID          uint    `json:"id_agen"`
	OutletName  string  `json:"nama_outlet"`
	Name        string  `json:"nama_agen"`
	PhoneNumber string  `json:"no_telp"`
	Province    string  `json:"provinsi"`
	City        string  `json:"kabko"`
	District    string  `json:"kecamatan"`
	Address     string  `json:"alamat_lengkap"`
	Rating      float64 `json:"rating"`
}

func (a *ListAgent) FromEntity(source entities.Agents) {
	helper.ConvertStruct(source, a)
	a.Province = source.Location.Province
	a.City = source.Location.City
	a.District = source.Location.District
	a.Address = source.Location.Address
}

type DetailAGent struct {
	ID          uint       `json:"id"`
	OutletName  string     `json:"nama_outlet"`
	Name        string     `json:"nama_agen"`
	PhoneNumber string     `json:"no_telp"`
	Province    string     `json:"provinsi"`
	City        string     `json:"kabko"`
	District    string     `json:"kecamatan"`
	Address     string     `json:"alamat_lengkap"`
	MaxTrx      int        `json:"maksimum_transaksi"`
	IsAvailable bool       `json:"status"`
	Rating      float64    `json:"rating"`
	Service     []*Service `json:"layanan"`
}

func (a *DetailAGent) FromEntity(source entities.Agents) {
	helper.ConvertStruct(source, a)
	a.Province = source.Location.Province
	a.City = source.Location.City
	a.District = source.Location.District
	a.Address = source.Location.Address
}

func (a *DetailAGent) ToEntity(Login, Location uint, service []*entities.Services) *entities.Agents {
	ent := entities.Agents{}
	helper.ConvertStruct(a, ent)
	ent.Name = a.Name
	ent.OutletName = a.OutletName
	ent.PhoneNumber = a.PhoneNumber
	ent.IsAvailable = a.IsAvailable
	ent.MaxTrx = a.MaxTrx
	ent.LocationID = Location
	ent.LoginID = Login
	ent.Services = service
	return &ent
}

func (a *DetailAGent) ToLocation() *Location {
	ent := Location{}
	helper.ConvertStruct(a, ent)
	ent.Province = a.Province
	ent.City = a.City
	ent.District = a.District
	ent.Address = a.Address
	return &ent
}

type Service struct {
	ID          uint   `json:"id"`
	ServiceName string `json:"judul"`
	Description string `json:"deskripsi"`
}

func (s *Service) FromEntity(source entities.Services) {
	helper.ConvertStruct(source, s)
}
