package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
	//"github.com/gofiber/fiber/v2"
	//"go/constant"
)

type ServiceTrasactionInterface interface {
}

type ServiceTrasaction struct {
	Repository repository.TransactionRepositoryInterface
}

func (S *ServiceTrasaction) Create(trans *contract.Transaction) (*contract.Transaction, error) {
	ent := trans.ToEntity()
	println(ent.Amount)
	println(ent.Tipe)

	agent := ServiceAgent{Repository: repository.AgentRepositoryMysql{}}
	cust := ServiceCustomer{Repository: repository.CustomerRepositoryMysql{}}

	agentEnt, agentErr := agent.Repository.GetByID(ent.AgentsID)
	custEnt, custErr := cust.Repository.GetByID(ent.CustomersID)

	if agentErr != nil || custErr != nil {
		return nil, agentErr
	}

	ent.Agents = *agentEnt
	ent.Customers = *custEnt
	ent.Location.Login = ent.Customers.Login

	res, err := S.Repository.Create(ent)
	if err != nil {
		return nil, err
	}
	contract := contract.Transaction{}
	contract.FromEntity(res)
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
		contract.FromEntity(v)
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
		contract.FromEntity(v)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}

func (S *ServiceTrasaction) GetByAgentID(id uint) ([]*contract.Transaction, error) {
	res, err := S.Repository.GetAllByID("agent", id)
	if err != nil {
		return nil, err
	}
	var contractList []*contract.Transaction
	for _, v := range res {
		contract := contract.Transaction{}
		contract.FromEntity(v)
		contractList = append(contractList, &contract)
	}
	return contractList, nil
}

func (S *ServiceTrasaction) SetStatus(id uint, status int) (*contract.Transaction, error) {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	res.Status = status

	res, errr := S.Repository.Update(res)
	if errr != nil {
		return nil, errr
	}
	contract := contract.Transaction{}
	helper.ConvertStruct(res, &contract)
	return &contract, nil
}

func (S *ServiceTrasaction) Delete(id uint) error {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return err
	}
	res, errr := S.Repository.Delete(res)
	if errr != nil {
		return errr
	}
	contract := contract.Transaction{}
	helper.ConvertStruct(res, &contract)
	return nil
}
