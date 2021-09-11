package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"

	"github.com/ulule/deepcopier"
	//"github.com/gofiber/fiber/v2"
	//"go/constant"
)

type ServiceTrasactionInterface interface {
}

type ServiceTrasaction struct {
	Repository repository.TransactionRepositoryInterface
}

func (S *ServiceTrasaction) Create(trans *contract.Transaction) (*contract.Transaction, error) {
	ent := entities.Transaction{}
	helper.ConvertStruct(trans, &ent)

	agent := ServiceAgent{Repository: repository.AgentRepositoryMysql{}}
	cust := ServiceCustomer{Repository: repository.CustomerRepositoryMysql{}}
	loc := ServiceLocation{Repository: repository.LocationRepositoryMysql{}}

	agentEnt, agentErr := agent.Repository.GetByID(ent.AgentsID)
	custEnt, custErr := cust.Repository.GetByID(ent.CustomersID)
	locEnt, locErr := loc.Repository.GetByCity(trans.City)
	if agentErr != nil || custErr != nil || locErr != nil {
		return nil, agentErr
	}

	ent.Agents = *agentEnt
	ent.Customers = *custEnt
	ent.Location = *locEnt

	res, err := S.Repository.Create(&ent)
	if err != nil {
		return nil, err
	}
	contract := contract.Transaction{}
	helper.ConvertStruct(res, &contract)
	return &contract, nil
}

func (S *ServiceTrasaction) GetAll() ([]*contract.Transaction, error) {
	res, err := S.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	var contractList []*contract.Transaction
	for _, v := range res {
		contract := contract.Transaction{}
		deepcopier.Copy(v).To(&contract)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}

func (S *ServiceTrasaction) GetByCustID(id uint) ([]*contract.Transaction, error) {
	res, err := S.Repository.GetAllByID("cust", id)
	if err != nil {
		return nil, err
	}
	var contractList []*contract.Transaction
	for _, v := range res {
		contract := contract.Transaction{}
		deepcopier.Copy(v).To(&contract)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}
