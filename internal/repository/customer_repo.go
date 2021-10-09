package repository

import (
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	// Create(user *entities.Customers) (*entities.Customers, error) //1
	GetByID(id int) (*entities.Customers, error) //1
	// Update() (*entities.Transaction, error)                      //1
	// Delete() error                                               //10
}

type CustomerRepositoryMysql struct {
	Db *gorm.DB
}

func (usr CustomerRepositoryMysql) GetByID(id uint) (*entities.Customers, error) {
	ent := entities.Customers{}
	err := usr.Db.Preload("Login").First(&ent, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}
