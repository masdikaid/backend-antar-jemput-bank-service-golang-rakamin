package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/entities"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
	"fmt"

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
	agentRepo := repository.AgentRepositoryMysql{}
	agent, errr := agentRepo.GetByID(ent.AgentsID)
	if errr != nil {
		return nil, errr
	}
	customerRepo := repository.CustomerRepositoryMysql{}
	customer, errr := customerRepo.GetByID(ent.CustomersID)
	if errr != nil {
		return nil, errr
	}
	locationRepo := repository.LocationRepositoryMysql{}
	location, errr := locationRepo.GetByCity(trans.City)
	if errr != nil {
		return nil, errr
	}
	ent.Agents = *agent
	ent.Customers = *customer
	ent.Location = *location

	res, err := S.Repository.Create(&ent)
	if err != nil {
		return nil, err
	}
	contract := contract.Transaction{}
	helper.ConvertStruct(res, &contract)
	fmt.Println(contract.Location.City)
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
