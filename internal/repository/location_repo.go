package repository

import (
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
)

type LocationRepositoryInterface interface {
	GetByCity(city string) (*entities.Location, error)
	GetByAgentID(agentID uint) (*entities.Location, error)
	Update(location *entities.Location) (*entities.Location, error)
}

type LocationRepositoryMysql struct {
	Db *gorm.DB
}

func NewLocationRepo(db *gorm.DB) LocationRepositoryInterface {
	return &LocationRepositoryMysql{Db: db}
}

func (usr LocationRepositoryMysql) GetByCity(city string) (*entities.Location, error) {
	ent := entities.Location{City: city}
	err := usr.Db.First(&ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}

func (usr LocationRepositoryMysql) GetByAgentID(agentID uint) (*entities.Location, error) {
	temp := entities.Agents{}
	err := usr.Db.First(&temp, agentID)
	if err.Error != nil {
		return nil, err.Error
	}

	ent := entities.Location{}
	errL := usr.Db.First(&ent, temp.LocationID)
	if errL.Error != nil {
		return nil, errL.Error
	}
	return &ent, nil
}

func (usr LocationRepositoryMysql) Update(location *entities.Location) (*entities.Location, error) {
	res := usr.Db.Save(location)
	if res.Error != nil {
		return nil, res.Error
	}
	return location, nil
}
