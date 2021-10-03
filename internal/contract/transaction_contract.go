package contract

import (
	"backend-a-antar-jemput/internal/entities"
	"time"
)

type Transaction struct {
	CustomersID uint   `json:"id_cust"`
	Services    string `json:"jenis_layanan"`
	Type        string `json:"jenis_transaksi"`
	Amount      uint   `json:"nominal_transaksi_idr"`
	Province    string `json:"provinsi"`
	City        string `json:"kabko"`
	District    string `json:"kecamatan"`
	Address     string `json:"alamat_lengkap"`
	AgentsID    uint   `json:"id_agen"`
}

func (t *Transaction) FromEntity(s *entities.Transaction) {
	t.CustomersID = s.CustomersID
	t.Services = s.Services.ServiceName
	t.Type = s.Type
	t.Amount = s.Amount
	t.Province = s.Province
	t.City = s.City
	t.District = s.District
	t.Address = s.Address
	t.AgentsID = s.AgentsID
}

func (t *Transaction) ToEntity() *entities.Transaction {
	ent := entities.Transaction{}
	ent.CustomersID = t.CustomersID
	ent.Services = entities.Services{ServiceName: t.Services}
	ent.Type = t.Type
	ent.Amount = t.Amount
	ent.Province = t.Province
	ent.City = t.City
	ent.District = t.District
	ent.Address = t.Address
	ent.AgentsID = t.AgentsID
	return &ent
}

type TransactionResponse struct {
	ID       uint             `json:"id"`
	Status   uint             `json:"status"`
	CreateAt time.Time        `json:"create_at"`
	Service  string           `json:"jenis_layanan"`
	Type     string           `json:"jenis_transaksi"`
	Amount   uint             `json:"nominal_transaksi"`
	Ratting  float64          `json:"rating"`
	Agent    ListAgent        `json:"agen"`
	Customer CustomerResponse `json:"customer"`
}

func (t *TransactionResponse) FromEntity(s *entities.Transaction) {
	agent := ListAgent{
		ID:          s.AgentsID,
		OutletName:  s.Agents.OutletName,
		Name:        s.Agents.Name,
		PhoneNumber: s.Agents.PhoneNumber,
		Rating:      s.Agents.Rating,
		Province:    s.Agents.Location.Province,
		City:        s.Agents.Location.City,
		District:    s.Agents.Location.District,
		Address:     s.Agents.Location.Address,
	}

	cust := CustomerResponse{
		ID:          s.Customers.ID,
		Name:        s.Customers.Name,
		PhoneNumber: s.Customers.PhoneNumber,
		Province:    s.Province,
		City:        s.City,
		District:    s.District,
		Address:     s.Address,
	}

	t.ID = s.ID
	t.Status = s.Status
	t.CreateAt = s.CreatedAt
	t.Service = s.Services.ServiceName
	t.Type = s.Type
	t.Amount = s.Amount
	t.Ratting = s.Rating
	t.Agent = agent
	t.Customer = cust
}

func (t *TransactionResponse) ToEntity() *entities.Transaction {
	ent := entities.Transaction{}
	ent.ID = t.ID
	ent.Status = t.Status
	ent.CreatedAt = t.CreateAt
	ent.Services = entities.Services{ServiceName: t.Service}
	ent.Type = t.Type
	ent.Amount = t.Amount
	ent.Rating = t.Ratting
	ent.AgentsID = t.Agent.ID
	ent.CustomersID = t.Customer.ID
	return &ent
}
