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
	Province    string  `json:"alamat_agen_provinsi"`
	City        string  `json:"alamat_agen_kabko"`
	District    string  `json:"alamat_agen_kecamatan"`
	Address     string  `json:"alamat_agen_lengkap"`
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
	Province    string     `json:"alamat_agen_provinsi"`
	City        string     `json:"alamat_agen_kabko"`
	District    string     `json:"alamat_agen_kecamatan"`
	Address     string     `json:"alamat_agen_lengkap"`
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

type Service struct {
	ID              uint   `json:"id"`
	ServiceName     string `json:"judul"`
	TransactionName string `json:"deskripsi"`
}

func (s *Service) FromEntity(source entities.Services) {
	helper.ConvertStruct(source, s)
}
