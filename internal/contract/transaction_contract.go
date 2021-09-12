package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/tools/helper"
)

// type Transaction struct{
// 	Tipe string `json:"jenis_transaksi"`
// 	Amount int `json:"nominal_transaksi_idr"`
// 	Status int `json:"Status"`
// 	//Rating float64 `json:Rating`
// 	//id int
// 	//customerId customer
// 	//agetId Agent
// 	//locationID Location
// }

type Transaction struct {
	Status      int     `json:"status"`
	CustomersID uint    `json:"id_cust"`
	Tipe        string  `json:"jenis_transaksi"`
	Amount      int     `json:"nominal_transaksi_idr"`
	Province    string  `json:"alamat_cust_provinsi"`
	City        string  `json:"alamat_cust_kabko"`
	District    string  `json:"alamat_cust_kecamatan"`
	Address     string  `json:"alamat_cust_lengkap"`
	AgentsID    uint    `json:"id_agen"`
	ID          uint    `json:"id_transaksi"`
	Rating      float64 `json:"rating"`
}

func (t *Transaction) FromEntity(source *entities.Transaction) {
	helper.ConvertStruct(source, t)
	t.Province = source.Location.Province
	t.City = source.Location.City
	t.District = source.Location.District
	t.Address = source.Location.Address
}

func (t *Transaction) ToEntity() *entities.Transaction {
	ent := entities.Transaction{}
	helper.ConvertStruct(t, ent)
	ent.Amount = t.Amount
	ent.Tipe = t.Tipe
	ent.Location.Province = t.Province
	ent.Location.City = t.City
	ent.Location.District = t.District
	ent.Location.Address = t.Address
	return &ent
}

// "jenis_transaksi": "Setoran Pinjaman",
// "nominal_transaksi_idr":1000000,
// "alamat_cust_provinsi": "JAWA BARAT",
// "alamat_cust_kabko": "BANDUNG",
// "alamat_cust_kecamatan":"SOREANG",
// "alamat_cust_lengkap":"Jl. Soreang No.181",
// "id_agen":3}

// type TransactionResponse struct{
// 	Status int `json:status`
// 	Message string `json:message`
// }
