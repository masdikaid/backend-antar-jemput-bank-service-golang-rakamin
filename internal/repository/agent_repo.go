package repository

import (
	"backend-a-antar-jemput/internal/entities"

	"gorm.io/gorm"
)

type AgentRepositoryInterface interface {
	GetByID(id uint) (*entities.Agents, error)
	Update(ent *entities.Agents) (*entities.Agents, error)
	GetAvailableAgent(service, district string, trx int) (*[]entities.Agents, error)
	GetServiceByID(id uint) (*[]entities.Services, error)
	// Create(user *entities.Agents) (*entities.Agents, error)
	// Delete() error
}

type AgentRepositoryMysql struct {
	Db *gorm.DB
}

func NewAgentRepo(db *gorm.DB) AgentRepositoryInterface {
	return &AgentRepositoryMysql{Db: db}
}

func (usr AgentRepositoryMysql) GetByID(id uint) (*entities.Agents, error) {
	ent := entities.Agents{}

	err := usr.Db.Preload("Login").Preload("Location").Preload("Services").First(&ent, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}

func (usr AgentRepositoryMysql) GetServiceByID(id uint) (*[]entities.Services, error) {
	ent := []entities.Services{}

	err := usr.Db.Joins("JOIN agent_services ON services.id=agent_services.services_id AND agent_services.agents_id=?", id).
		Find(&ent)
	if err.Error != nil {
		return nil, err.Error
	}

	return &ent, nil
}

func (usr AgentRepositoryMysql) GetAvailableAgent(service, district string, trx int) (*[]entities.Agents, error) {
	ent := []entities.Agents{}
	ser := entities.Services{}

	err := usr.Db.Where("service_name=?", service).First(&ser)
	if err.Error != nil {
		return nil, err.Error
	}

	errr := usr.Db.Preload("Location").Joins("Location").Joins("JOIN agent_services ON agents.id=agent_services.agents_id AND agent_services.services_id=?", ser.ID).
		Where("district=? AND max_trx>=? AND is_available=?", district, trx, true).Find(&ent)
	if errr.Error != nil {
		return nil, errr.Error
	}

	return &ent, nil
}

func (usr AgentRepositoryMysql) Update(ent *entities.Agents) (*entities.Agents, error) {
	err := usr.Db.Save(ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return ent, nil
}

// func (usr AgentRepositoryMysql) Create(ent *entities.Agents) (*entities.Agents, error) {
//
// 	res := databases.Create(&ent)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return ent, nil
// }
