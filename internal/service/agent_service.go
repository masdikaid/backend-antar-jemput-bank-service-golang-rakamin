package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
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
	contract.FromEntity(*res)
	return &contract, nil
}

func (S ServiceAgent) GetListAgent(district string, trx int) ([]*contract.Agent, error) {
	res, err := S.Repository.GetAvailableAgent(district, trx)
	if err != nil {
		return nil, err
	}

	var contractList []*contract.Agent
	for _, v := range *res {
		contract := contract.Agent{}
		contract.FromEntity(v)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}

func (S ServiceAgent) UpdateRating(id uint, rating int) error {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return err
	}

	if rating == 0 {
		return nil
	}
	res.Rating = (res.Rating + float32(rating)) / 2

	_, errr := S.Repository.Update(res)
	if errr != nil {
		return errr
	}
	return nil
}
