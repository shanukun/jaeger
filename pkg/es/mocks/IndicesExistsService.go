// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IndicesExistsService is an autogenerated mock type for the IndicesExistsService type
type IndicesExistsService struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx
func (_m *IndicesExistsService) Do(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIndicesExistsService creates a new instance of IndicesExistsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndicesExistsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndicesExistsService {
	mock := &IndicesExistsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
