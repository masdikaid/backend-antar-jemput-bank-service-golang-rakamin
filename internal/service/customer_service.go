package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
)

type ServiceCustomerInterface interface {
}

type ServiceCustomer struct {
	Repository repository.CustomerRepositoryInterface
}

func (S ServiceCustomer) GetCustomer(id uint) (*contract.Customer, error) {
	res, err := S.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	contract := contract.Customer{}
	helper.ConvertStruct(res, &contract)
	return &contract, nil
}
