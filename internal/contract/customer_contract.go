package contract

type Customer struct {
	CustomerID int          `json:"id_cust"`
	Service    string       `json:"jenis_layanan"`
	Tipe       string       `json:"jenis_transaksi"`
	Amount     int          `json:"nominal_transaksi_idr"`
	Province   string       `json:"provinsi"`
	City       string       `json:"kabko"`
	District   string       `json:"kecamatan"`
	Address    string       `json:"alamat_lengkap"`
	ListAgent  []*ListAgent `json:"list_rekomendasi_agen"`
}

type CustomerResponse struct {
	ID          uint   `json:"id_cust"`
	Name        string `json:"nama"`
	PhoneNumber string `json:"no_telp"`
	Province    string `json:"provinsi"`
	City        string `json:"kabko"`
	District    string `json:"kecamatan"`
	Address     string `json:"alamat_lengkap"`
}
