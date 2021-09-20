package repository

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
)

type AgentRepositoryInterface interface {
	// Create(user *entities.Agents) (*entities.Agents, error) //1
	GetByID(id uint) (entities.Agents, error) //1
	// Update() (*entities.Transaction, error)                      //1
	// Delete() error                                               //10

	GetAvailableAgent(district string, trx int) (entities.Agents, error)
}

type AgentRepositoryMysql struct {
}

func (usr AgentRepositoryMysql) GetByID(id uint) (*entities.Agents, error) {
	ent := entities.Agents{}
	databases.Load()
	err := databases.DBCon.Preload("Login").First(&ent, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}

func (usr AgentRepositoryMysql) GetAvailableAgent(district string, trx int) (*[]entities.Agents, error) {
	ent := []entities.Agents{}
	databases.Load()

	err := databases.DBCon.Preload("Location").Joins("Location").
		Where("district=? AND max_trx>=? AND is_available=?", district, trx, true).Find(&ent)
	if err.Error != nil {
		return nil, err.Error
	}

	// if len(ent) == 0 {
	// 	err := databases.DBCon.Preload("Location").Joins("Location").
	// 		Where("city=? AND max_trx>=? AND is_available=?", city, trx, true).Find(&ent)
	// 	if err.Error != nil {
	// 		return nil, err.Error
	// 	}

	// 	return &ent, nil
	// }

	return &ent, nil
}
