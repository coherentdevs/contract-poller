// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	raw "github.com/coherent-api/contract-service/protos/go/protos/evm/raw"
	mock "github.com/stretchr/testify/mock"
)

// Streamer is an autogenerated mock type for the Streamer type
type Streamer struct {
	mock.Mock
}

// Stream provides a mock function with given fields: data
func (_m *Streamer) Stream(data *raw.Data) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*raw.Data) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStreamer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStreamer creates a new instance of Streamer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStreamer(t mockConstructorTestingTNewStreamer) *Streamer {
	mock := &Streamer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
