package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
)

type ServiceLocationInterface interface {
}

type ServiceLocation struct {
	Repository repository.LocationRepositoryMysql
}

func (S ServiceLocation) GetLocation(city string) (*contract.Location, error) {
	res, err := S.Repository.GetByCity(city)
	if err != nil {
		return nil, err
	}

	contract := contract.Location{}
	helper.ConvertStruct(res, &contract)
	return &contract, nil
}
