package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
)

type ServiceAgentInterface interface {
}

type ServiceAgent struct {
	Repository repository.AgentRepositoryInterface
}

func (S ServiceAgent) GetAgent(id uint) (*contract.DetailAGent, error) {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	contract := contract.DetailAGent{}
	contract.FromEntity(*res)
	return &contract, nil
}

func (S ServiceAgent) GetAgentService(id uint) ([]*contract.Service, error) {
	res, err := S.Repository.GetServiceByID(id)
	if err != nil {
		return nil, err
	}
	var contractList []*contract.Service
	for _, v := range *res {
		contract := contract.Service{}
		contract.FromEntity(v)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}

func (S ServiceAgent) GetListAgent(service, district string, trx int) ([]*contract.ListAgent, error) {
	res, err := S.Repository.GetAvailableAgent(service, district, trx)
	if err != nil {
		return nil, err
	}

	var contractList []*contract.ListAgent
	for _, v := range *res {
		contract := contract.ListAgent{}
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

	if res.Rating == 0 {
		res.Rating = float64(rating)
		_, errr := S.Repository.Update(res)
		if errr != nil {
			return errr
		}
		return nil
	}

	if rating == 0 {
		return nil
	}
	res.Rating = (res.Rating + float64(rating)) / 2

	_, errr := S.Repository.Update(res)
	if errr != nil {
		return errr
	}
	return nil
}

// func (S ServiceAgent) Create(agent *contract.DetailAGent) (*contract.Agent, error) {
// 	ent := agent.ToEntity()

// 	agent := ServiceAgent{Repository: repository.AgentRepositoryMysql{}}
// 	agentL, agentErr := agent.Repository.GetByID(ent.AgentsID)
// 	if agentErr != nil || custErr != nil {
// 		return nil, agentErr
// 	}

// 	ent.Login=*
// 	ent.Location.Login.ID = agent.LoginID

// 	res, err := S.Repository.Create(ent)
// 	if err != nil {
// 		return nil, err
// 	}
// 	contract := contract.Agent{}
// 	contract.FromEntity(*res)
// 	return &contract, nil
// }
