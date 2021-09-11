package repository

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
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
	ent := entities.Agents{Users: entities.Users{Model: gorm.Model{ID: id}}}
	databases.Load()
	err := databases.DBCon.First(&ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}
