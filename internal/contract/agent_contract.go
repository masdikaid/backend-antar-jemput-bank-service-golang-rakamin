package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/tools/helper"
)

type Agent struct {
	ID          uint   `json:"id_agen"`
	Name        string `json:"nama_agen"`
	OutletName  string `json:"nama_outlet"`
	Address     string `json:"alamat_agen_lengkap"`
	District    string `json:"alamat_agen_kecamatan"`
	City        string `json:"alamat_agen_kabko"`
	Province    string `json:"alamat_agen_provinsi"`
	PhoneNumber string `json:"no_wa"`
}

func (a *Agent) FromEntity(source entities.Agents) {
	helper.ConvertStruct(source, a)
	a.Province = source.Location.Province
	a.City = source.Location.City
	a.District = source.Location.District
	a.Address = source.Location.Address
}
