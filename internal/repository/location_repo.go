package repository

import (
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
)

type LocationRepositoryInterface interface {
	// Create(user *entities.Location) (*entities.Location, error) //1
	GetByID(id uint) (entities.Location, error) //1
	// Update() (*entities.Transaction, error)                      //1
	// Delete() error                                               //10
}

type LocationRepositoryMysql struct {
	Db *gorm.DB
}

func (usr LocationRepositoryMysql) GetByCity(city string) (*entities.Location, error) {
	ent := entities.Location{City: city}
	err := usr.Db.First(&ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}
