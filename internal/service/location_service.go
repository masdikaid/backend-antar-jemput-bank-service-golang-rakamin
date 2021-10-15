package service

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/internal/repository"
	"backend-a-antar-jemput/tools/helper"
)

type ServiceLocationInterface interface {
	GetLocation(city string) (*contract.Location, error)
	Update(loc *contract.Location, agentID uint) error
}

type ServiceLocation struct {
	Repository repository.LocationRepositoryInterface
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

func (S ServiceLocation) Update(loc *contract.Location, agentID uint) error {
	temp, err := S.Repository.GetByAgentID(agentID)
	if err != nil {
		return err
	}

	temp.Province = loc.Province
	temp.City = loc.City
	temp.District = loc.District
	temp.Address = loc.Address

	_, errL := S.Repository.Update(temp)
	if errL != nil {
		return errL
	}

	return nil
}
