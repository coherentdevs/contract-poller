// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	constants "github.com/coherent-api/contract-poller/shared/constants"

	mock "github.com/stretchr/testify/mock"

	models "github.com/coherent-api/contract-poller/poller/pkg/models"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// DeleteContractByAddress provides a mock function with given fields: address
func (_m *Database) DeleteContractByAddress(address string) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEventFragment provides a mock function with given fields: eventFragment
func (_m *Database) DeleteEventFragment(eventFragment *models.EventFragment) error {
	ret := _m.Called(eventFragment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.EventFragment) error); ok {
		r0 = rf(eventFragment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMethodFragment provides a mock function with given fields: methodFragment
func (_m *Database) DeleteMethodFragment(methodFragment *models.MethodFragment) error {
	ret := _m.Called(methodFragment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MethodFragment) error); ok {
		r0 = rf(methodFragment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetContract provides a mock function with given fields: contractAddress, blockchain
func (_m *Database) GetContract(contractAddress string, blockchain constants.Blockchain) (*models.Contract, error) {
	ret := _m.Called(contractAddress, blockchain)

	var r0 *models.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func(string, constants.Blockchain) (*models.Contract, error)); ok {
		return rf(contractAddress, blockchain)
	}
	if rf, ok := ret.Get(0).(func(string, constants.Blockchain) *models.Contract); ok {
		r0 = rf(contractAddress, blockchain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func(string, constants.Blockchain) error); ok {
		r1 = rf(contractAddress, blockchain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractsToBackfill provides a mock function with given fields:
func (_m *Database) GetContractsToBackfill() ([]models.Contract, error) {
	ret := _m.Called()

	var r0 []models.Contract
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Contract, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Contract); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Contract)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventFragmentById provides a mock function with given fields: eventId
func (_m *Database) GetEventFragmentById(eventId string) (*models.EventFragment, error) {
	ret := _m.Called(eventId)

	var r0 *models.EventFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.EventFragment, error)); ok {
		return rf(eventId)
	}
	if rf, ok := ret.Get(0).(func(string) *models.EventFragment); ok {
		r0 = rf(eventId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.EventFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(eventId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMethodFragmentByID provides a mock function with given fields: methodId
func (_m *Database) GetMethodFragmentByID(methodId string) (*models.MethodFragment, error) {
	ret := _m.Called(methodId)

	var r0 *models.MethodFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.MethodFragment, error)); ok {
		return rf(methodId)
	}
	if rf, ok := ret.Get(0).(func(string) *models.MethodFragment); ok {
		r0 = rf(methodId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MethodFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(methodId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContractByAddress provides a mock function with given fields: contract
func (_m *Database) UpdateContractByAddress(contract *models.Contract) error {
	ret := _m.Called(contract)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Contract) error); ok {
		r0 = rf(contract)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEventFragment provides a mock function with given fields: eventFragment
func (_m *Database) UpdateEventFragment(eventFragment *models.EventFragment) error {
	ret := _m.Called(eventFragment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.EventFragment) error); ok {
		r0 = rf(eventFragment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMethodFragment provides a mock function with given fields: methodFragment
func (_m *Database) UpdateMethodFragment(methodFragment *models.MethodFragment) error {
	ret := _m.Called(methodFragment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MethodFragment) error); ok {
		r0 = rf(methodFragment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertContracts provides a mock function with given fields: contracts
func (_m *Database) UpsertContracts(contracts []models.Contract) (int64, error) {
	ret := _m.Called(contracts)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func([]models.Contract) (int64, error)); ok {
		return rf(contracts)
	}
	if rf, ok := ret.Get(0).(func([]models.Contract) int64); ok {
		r0 = rf(contracts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func([]models.Contract) error); ok {
		r1 = rf(contracts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertEventFragment provides a mock function with given fields: eventFragment
func (_m *Database) UpsertEventFragment(eventFragment *models.EventFragment) (int64, error) {
	ret := _m.Called(eventFragment)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.EventFragment) (int64, error)); ok {
		return rf(eventFragment)
	}
	if rf, ok := ret.Get(0).(func(*models.EventFragment) int64); ok {
		r0 = rf(eventFragment)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*models.EventFragment) error); ok {
		r1 = rf(eventFragment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertMethodFragment provides a mock function with given fields: methodFragment
func (_m *Database) UpsertMethodFragment(methodFragment *models.MethodFragment) (int64, error) {
	ret := _m.Called(methodFragment)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.MethodFragment) (int64, error)); ok {
		return rf(methodFragment)
	}
	if rf, ok := ret.Get(0).(func(*models.MethodFragment) int64); ok {
		r0 = rf(methodFragment)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*models.MethodFragment) error); ok {
		r1 = rf(methodFragment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
