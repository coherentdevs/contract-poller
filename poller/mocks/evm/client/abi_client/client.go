// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	constants "github.com/datadaodevs/go-service-framework/constants"

	etherscan "github.com/nanmu42/etherscan-api"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ContractSource provides a mock function with given fields: ctx, contractAddress, blockchain
func (_m *Client) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	ret := _m.Called(ctx, contractAddress, blockchain)

	var r0 etherscan.ContractSource
	if rf, ok := ret.Get(0).(func(context.Context, string, constants.Blockchain) etherscan.ContractSource); ok {
		r0 = rf(ctx, contractAddress, blockchain)
	} else {
		r0 = ret.Get(0).(etherscan.ContractSource)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, constants.Blockchain) error); ok {
		r1 = rf(ctx, contractAddress, blockchain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
