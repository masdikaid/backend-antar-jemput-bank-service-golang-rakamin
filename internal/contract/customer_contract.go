package contract

type Customer struct {
	CustomerID int          `json:"id_cust"`
	Service    string       `json:"jenis_layanan"`
	Tipe       string       `json:"jenis_transaksi"`
	Amount     int          `json:"nominal_transaksi_idr"`
	Province   string       `json:"alamat_cust_provinsi"`
	City       string       `json:"alamat_cust_kabko"`
	District   string       `json:"alamat_cust_kecamatan"`
	Address    string       `json:"alamat_cust_lengkap"`
	ListAgent  []*ListAgent `json:"list_rekomendasi_agen"`
}
