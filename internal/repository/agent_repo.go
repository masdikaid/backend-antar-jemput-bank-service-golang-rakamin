package repository

import (
	"backend-a-antar-jemput/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

type AgentRepositoryInterface interface {
	GetByID(id uint) (*entities.Agents, error)
	GetServiceByName(name string) (*entities.Services, error)
	GetServiceByID(id uint) (*[]entities.Services, error)
	GetAvailableAgent(service, city, district string, trx int) (*[]entities.Agents, error)
	Update(ent *entities.Agents) (*entities.Agents, error)
}

type AgentRepositoryMysql struct {
	Db *gorm.DB
}

func NewAgentRepo(db *gorm.DB) AgentRepositoryInterface {
	return &AgentRepositoryMysql{Db: db}
}

func (usr AgentRepositoryMysql) GetByID(id uint) (*entities.Agents, error) {
	ent := entities.Agents{}

	err := usr.Db.Preload("Login").Preload("Location").First(&ent, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &ent, nil
}

func (usr AgentRepositoryMysql) GetServiceByName(name string) (*entities.Services, error) {
	ent := entities.Services{}

	err := usr.Db.First(&ent, "service_name=?", name)
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

func (usr AgentRepositoryMysql) GetAvailableAgent(service, city, district string, trx int) (*[]entities.Agents, error) {
	ent := []entities.Agents{}
	ser := entities.Services{}

	err := usr.Db.Where("service_name=?", service).First(&ser)
	if err.Error != nil {
		return nil, err.Error
	}

	// errr := usr.Db.Preload("Location").Joins("Location").Joins("JOIN agent_services ON agents.id=agent_services.agents_id AND agent_services.services_id=?", ser.ID).
	// 	Where("city=? AND district=? AND max_trx>=? AND is_available=?", city, district, trx, true).Find(&ent)
	// if errr.Error != nil {
	// 	return nil, errr.Error
	// }

	errr := usr.Db.Preload("Location").Find(&ent)
	if errr.Error != nil {
		return nil, errr.Error
	}
	fmt.Println(ent)

	fmt.Println(service, city, district, trx)
	return &ent, nil
}

func (usr AgentRepositoryMysql) Update(ent *entities.Agents) (*entities.Agents, error) {
	errS := usr.Db.Table("agent_services").Where("agents_id=?", ent.ID).Delete("agent_services")
	if errS.Error != nil {
		return nil, errS.Error
	}

	err := usr.Db.Save(ent)
	if err.Error != nil {
		return nil, err.Error
	}

	resp := entities.Agents{}
	errR := usr.Db.Preload("Location").First(&resp, ent.ID)
	if errR.Error != nil {
		return nil, err.Error
	}
	return &resp, nil
}
