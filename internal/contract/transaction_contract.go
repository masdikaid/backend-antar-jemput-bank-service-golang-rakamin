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

type TransactionResponse struct{
	Status int `json:status`
	Message string `json:message`
}