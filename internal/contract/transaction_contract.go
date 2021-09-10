package contract

type Transaction struct{
	Tipe string `json:"jenis_transaksi"`
	Amount int `json:"nominal_transaksi_idr"`
	Status int `json:"Status"`
	//Rating float64 `json:Rating`
	//id int
	//customerId customer
	//agetId Agent
	//locationID Location
}

type TransactionCust struct{
	Status int `json:"Status"`
	Tipe string `json:"Tipe"`
	Amount int `json:"nominal_transaksi_idr"`
	LocationID uint `json:"Id Lokasi"`
	//Location Location
	AgentsID uint `json:"Id Agent"`
	Id uint `json:"Id Transaksi"`
	CustomersID uint `json:"Id Customer"`
	//Rating float64
}



// type TransactionResponse struct{
// 	Status int `json:status`
// 	Message string `json:message`
// }