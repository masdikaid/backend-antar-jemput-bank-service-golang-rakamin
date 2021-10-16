package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/internal/repository"
)

type ServiceAgentInterface interface {
	GetAgent(id uint) (*contract.DetailAGent, error)
	GetAgentService(id uint) ([]*contract.Service, error)
	GetListAgent(service, city, district string, trx int) ([]*contract.ListAgent, error)
	UpdateRating(id uint, rating int) error
	UpdateAgent(agent *contract.DetailAGent, agentID uint) (*contract.DetailAGent, error)
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

func (S ServiceAgent) GetListAgent(service, city, district string, trx int) ([]*contract.ListAgent, error) {
	res, err := S.Repository.GetAvailableAgent(service, city, district, trx)
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

func (S ServiceAgent) UpdateAgent(agent *contract.DetailAGent, agentID uint) (*contract.DetailAGent, error) {
	res, err := S.Repository.GetByID(agentID)
	if err != nil {
		return nil, err
	}
	listSer := []*entities.Services{}
	for _, ser := range agent.Service {
		temp, errS := S.Repository.GetServiceByName(ser.ServiceName)
		if errS != nil {
			return nil, errS
		}
		listSer = append(listSer, temp)
	}

	res.Name = agent.Name
	res.OutletName = agent.OutletName
	res.PhoneNumber = agent.PhoneNumber
	res.MaxTrx = agent.MaxTrx
	res.IsAvailable = agent.IsAvailable
	res.Services = listSer

	resA, errA := S.Repository.Update(res)
	if errA != nil {
		return nil, errA
	}
	contract := contract.DetailAGent{}
	contract.FromEntity(*resA)
	return &contract, nil
}
