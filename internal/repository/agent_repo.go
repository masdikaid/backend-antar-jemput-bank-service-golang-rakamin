package repository

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
)

type AgentRepositoryInterface interface {
	GetByID(id uint) (*entities.Agents, error)
	Update(ent *entities.Agents) (*entities.Agents, error)
	GetAvailableAgent(district string, trx int) (*[]entities.Agents, error)
	// Create(user *entities.Agents) (*entities.Agents, error)
	// Delete() error
}

type AgentRepositoryMysql struct {
}

func (usr AgentRepositoryMysql) GetByID(id uint) (*entities.Agents, error) {
	ent := entities.Agents{}
	databases.Load()
	err := databases.DBCon.Preload("Login").Preload("Location").Preload("Services").First(&ent, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}

func (usr AgentRepositoryMysql) GetServiceByID(id uint) (*[]entities.Services, error) {
	ent := []entities.Services{}
	databases.Load()

	err := databases.DBCon.Joins("JOIN agent_services ON services.id=agent_services.services_id AND agent_services.agents_id=?", id).
		Find(&ent)
	if err.Error != nil {
		return nil, err.Error
	}

	return &ent, nil
}

func (usr AgentRepositoryMysql) GetAvailableAgent(service, tipe, district string, trx int) (*[]entities.Agents, error) {
	ent := []entities.Agents{}
	ser := entities.Services{}
	databases.Load()
	err := databases.DBCon.Where("service_name=? AND transaction_name=?", service, tipe).First(&ser)
	if err.Error != nil {
		return nil, err.Error
	}

	errr := databases.DBCon.Preload("Location").Joins("Location").Joins("JOIN agent_services ON agents.id=agent_services.agents_id AND agent_services.services_id=?", ser.ID).
		Where("district=? AND max_trx>=? AND is_available=?", district, trx, true).Find(&ent)
	if errr.Error != nil {
		return nil, errr.Error
	}

	return &ent, nil
}

func (usr AgentRepositoryMysql) Update(ent *entities.Agents) (*entities.Agents, error) {
	databases.Load()
	err := databases.DBCon.Save(ent)
	if err.Error != nil {
		return nil, err.Error
	}
	return ent, nil
}

// func (usr AgentRepositoryMysql) Create(ent *entities.Agents) (*entities.Agents, error) {
// 	databases.Load()
// 	res := databases.DBCon.Create(&ent)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return ent, nil
// }
