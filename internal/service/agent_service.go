package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
)

type ServiceAgentInterface interface {
}

type ServiceAgent struct {
	Repository repository.AgentRepositoryMysql
}

func (S ServiceAgent) GetAgent(id uint) (*contract.Agent, error) {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	contract := contract.Agent{}
	helper.ConvertStruct(res, &contract)
	return &contract, nil
}
