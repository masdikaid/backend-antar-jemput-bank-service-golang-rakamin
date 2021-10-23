package repository

import (
	"backend-a-antar-jemput/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	Create(ent *entities.Transaction) (*entities.Transaction, error)  //1
	GetAll() ([]*entities.Transaction, error)                         //1
	GetAllByID(tipe string, id uint) ([]*entities.Transaction, error) //1
	Update(ent *entities.Transaction) (*entities.Transaction, error)
	GetByID(id uint) (*entities.Transaction, error)
	Delete(ent *entities.Transaction) (*entities.Transaction, error)
	// Update() (*entities.Transaction, error)                      //1
	// Delete() error                                               //10
}

type TransactionRepositoryMysql struct {
	Db *gorm.DB
}

func (T TransactionRepositoryMysql) Create(ent *entities.Transaction) (*entities.Transaction, error) {
	res := T.Db.Create(&ent)
	if res.Error != nil {
		return nil, res.Error
	}
	return ent, nil
}

func (T TransactionRepositoryMysql) GetAll() ([]*entities.Transaction, error) {
	res := []*entities.Transaction{}
	err := T.Db.Preload("Agents.Location").Preload("Customers").Order("id desc").Find(&res)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func (T TransactionRepositoryMysql) GetAllByID(tipe string, id uint) ([]*entities.Transaction, error) {
	var query string
	switch tipe {
	case "cust":
		query = "customers_id = ?"
	case "agent":
		query = "agents_id = ?"
	default:
		err := errors.New("wrong tipe")
		return nil, err
	}

	res := []*entities.Transaction{}
	err := T.Db.Preload("Agents.Location").Preload("Customers").Where(query, id).Order("id desc").Find(&res)
	if err.Error != nil {
		return nil, err.Error
	}
	return res, nil
}

func (T TransactionRepositoryMysql) Update(ent *entities.Transaction) (*entities.Transaction, error) {
	err := T.Db.Save(ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return ent, nil

}

func (T TransactionRepositoryMysql) Delete(ent *entities.Transaction) (*entities.Transaction, error) {
	err := T.Db.Delete(ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return ent, nil

}

func (T TransactionRepositoryMysql) GetByID(id uint) (*entities.Transaction, error) {
	ent := entities.Transaction{}
	err := T.Db.Preload("Agents.Location").Preload("Customers").Where("id = ?", id).First(&ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}
