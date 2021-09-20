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

	GetAvailableAgent(district, city string) (entities.Agents, error)
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

func (usr AgentRepositoryMysql) GetAvailableAgent(district, city string) (*[]entities.Agents, error) {
	ent := []entities.Agents{}
	databases.Load()

	err := databases.DBCon.Preload("Location").Joins("Location").Where("district=?", district).Find(&ent)
	if err.Error != nil {
		return nil, err.Error
	}

	// for _, v := range ent {
	// 	if v.ID == 0 {
	// 		err := databases.DBCon.Preload("Location").Joins("Location").Where("city=?", city).Find(&ent)
	// 		if err.Error != nil {
	// 			return nil, err.Error
	// 		}
	// 	}
	// }

	return &ent, nil
}

// databases.Load()
// 	res := []*entities.Transaction{}
// 	err := databases.DBCon.Preload("Location").Find(&res)
// 	if err.Error != nil {
// 		return nil, err.Error
// 	}
// 	return res, nil
