// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	helpers "user-service/internal/pkg/helpers"

	mock "github.com/stretchr/testify/mock"
)

// MongodbRepositoryQuery is an autogenerated mock type for the MongodbRepositoryQuery type
type MongodbRepositoryQuery struct {
	mock.Mock
}

// FindOneByEmail provides a mock function with given fields: ctx, email
func (_m *MongodbRepositoryQuery) FindOneByEmail(ctx context.Context, email string) <-chan helpers.Result {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindOneByEmail")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan helpers.Result); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindOneByEmailUserTemp provides a mock function with given fields: ctx, email
func (_m *MongodbRepositoryQuery) FindOneByEmailUserTemp(ctx context.Context, email string) <-chan helpers.Result {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindOneByEmailUserTemp")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan helpers.Result); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindOneUserId provides a mock function with given fields: ctx, userId
func (_m *MongodbRepositoryQuery) FindOneUserId(ctx context.Context, userId string) <-chan helpers.Result {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for FindOneUserId")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan helpers.Result); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// NewMongodbRepositoryQuery creates a new instance of MongodbRepositoryQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMongodbRepositoryQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MongodbRepositoryQuery {
	mock := &MongodbRepositoryQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
