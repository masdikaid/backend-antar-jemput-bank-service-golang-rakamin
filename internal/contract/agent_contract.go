package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/tools/helper"
)

type Agent struct {
	ID          uint    `json:"id_agen"`
	OutletName  string  `json:"nama_outlet"`
	Name        string  `json:"nama_agen"`
	PhoneNumber string  `json:"no_telp"`
	Province    string  `json:"alamat_agen_provinsi"`
	City        string  `json:"alamat_agen_kabko"`
	District    string  `json:"alamat_agen_kecamatan"`
	Address     string  `json:"alamat_agen_lengkap"`
	Rating      float32 `json:"rating"`
}

func (a *Agent) FromEntity(source entities.Agents) {
	helper.ConvertStruct(source, a)
	a.Province = source.Location.Province
	a.City = source.Location.City
	a.District = source.Location.District
	a.Address = source.Location.Address
}
