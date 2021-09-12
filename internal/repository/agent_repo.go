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
