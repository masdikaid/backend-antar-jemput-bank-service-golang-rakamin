package contract

import "backend-a-antar-jemput/internal/entities"

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
	Status      int    `json:"status"`
	CustomersID uint   `json:"id_cust"`
	Tipe        string `json:"jenis_transaksi"`
	Amount      int    `json:"nominal_transaksi_idr"`
	Province    string `json:"alamat_cust_provinsi"`
	City        string `json:"alamat_cust_kabko"`
	District    string `json:"alamat_cust_kecamatan"`
	Address     string `json:"alamat_cust_lengkap"`
	LocationID  uint   `json:"Id Lokasi"`
	Location    entities.Location
	AgentsID    uint    `json:"id_agen"`
	Id          uint    `json:"id_transaksi"`
	Rating      float64 `json:"rating"`
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
