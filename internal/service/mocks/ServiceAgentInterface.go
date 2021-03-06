// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contract "backend-a-antar-jemput/internal/contract"

	mock "github.com/stretchr/testify/mock"
)

// ServiceAgentInterface is an autogenerated mock type for the ServiceAgentInterface type
type ServiceAgentInterface struct {
	mock.Mock
}

// GetAgent provides a mock function with given fields: id
func (_m *ServiceAgentInterface) GetAgent(id uint) (*contract.DetailAGent, error) {
	ret := _m.Called(id)

	var r0 *contract.DetailAGent
	if rf, ok := ret.Get(0).(func(uint) *contract.DetailAGent); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.DetailAGent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgentService provides a mock function with given fields: id
func (_m *ServiceAgentInterface) GetAgentService(id uint) ([]*contract.Service, error) {
	ret := _m.Called(id)

	var r0 []*contract.Service
	if rf, ok := ret.Get(0).(func(uint) []*contract.Service); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contract.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListAgent provides a mock function with given fields: _a0, city, district, trx
func (_m *ServiceAgentInterface) GetListAgent(_a0 string, city string, district string, trx int) ([]*contract.ListAgent, error) {
	ret := _m.Called(_a0, city, district, trx)

	var r0 []*contract.ListAgent
	if rf, ok := ret.Get(0).(func(string, string, string, int) []*contract.ListAgent); ok {
		r0 = rf(_a0, city, district, trx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contract.ListAgent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int) error); ok {
		r1 = rf(_a0, city, district, trx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAgent provides a mock function with given fields: agent, agentID
func (_m *ServiceAgentInterface) UpdateAgent(agent *contract.DetailAGent, agentID uint) (*contract.DetailAGent, error) {
	ret := _m.Called(agent, agentID)

	var r0 *contract.DetailAGent
	if rf, ok := ret.Get(0).(func(*contract.DetailAGent, uint) *contract.DetailAGent); ok {
		r0 = rf(agent, agentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contract.DetailAGent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*contract.DetailAGent, uint) error); ok {
		r1 = rf(agent, agentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRating provides a mock function with given fields: id, rating
func (_m *ServiceAgentInterface) UpdateRating(id uint, rating int) error {
	ret := _m.Called(id, rating)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, int) error); ok {
		r0 = rf(id, rating)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
