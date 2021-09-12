package contract

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
	Province    string  `json:"alamat_cust_provinsi" deepcopier:"field:Location.Province"`
	City        string  `json:"alamat_cust_kabko" deepcopier:"field:Location.City"`
	District    string  `json:"alamat_cust_kecamatan" deepcopier:"field:Location.District"`
	Address     string  `json:"alamat_cust_lengkap" deepcopier:"field:Location.Address"`
	AgentsID    uint    `json:"id_agen"`
	ID          uint    `json:"id_transaksi"`
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
