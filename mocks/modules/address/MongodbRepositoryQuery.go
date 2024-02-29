// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	helpers "user-service/internal/pkg/helpers"

	mock "github.com/stretchr/testify/mock"

	request "user-service/internal/modules/address/models/request"
)

// MongodbRepositoryQuery is an autogenerated mock type for the MongodbRepositoryQuery type
type MongodbRepositoryQuery struct {
	mock.Mock
}

// FindCitiesByParam provides a mock function with given fields: ctx, payload
func (_m *MongodbRepositoryQuery) FindCitiesByParam(ctx context.Context, payload request.City) <-chan helpers.Result {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for FindCitiesByParam")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, request.City) <-chan helpers.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindContinent provides a mock function with given fields: ctx
func (_m *MongodbRepositoryQuery) FindContinent(ctx context.Context) <-chan helpers.Result {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindContinent")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context) <-chan helpers.Result); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindCountries provides a mock function with given fields: ctx, payload
func (_m *MongodbRepositoryQuery) FindCountries(ctx context.Context, payload request.Country) <-chan helpers.Result {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for FindCountries")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, request.Country) <-chan helpers.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindDistrictByParam provides a mock function with given fields: ctx, payload
func (_m *MongodbRepositoryQuery) FindDistrictByParam(ctx context.Context, payload request.District) <-chan helpers.Result {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for FindDistrictByParam")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, request.District) <-chan helpers.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindOneCountry provides a mock function with given fields: ctx, id
func (_m *MongodbRepositoryQuery) FindOneCountry(ctx context.Context, id int) <-chan helpers.Result {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindOneCountry")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, int) <-chan helpers.Result); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindOneSubdistrict provides a mock function with given fields: ctx, id
func (_m *MongodbRepositoryQuery) FindOneSubdistrict(ctx context.Context, id string) <-chan helpers.Result {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindOneSubdistrict")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, string) <-chan helpers.Result); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindProvinces provides a mock function with given fields: ctx, payload
func (_m *MongodbRepositoryQuery) FindProvinces(ctx context.Context, payload request.Province) <-chan helpers.Result {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for FindProvinces")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, request.Province) <-chan helpers.Result); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan helpers.Result)
		}
	}

	return r0
}

// FindSubDistrictByParam provides a mock function with given fields: ctx, payload
func (_m *MongodbRepositoryQuery) FindSubDistrictByParam(ctx context.Context, payload request.SubDistrict) <-chan helpers.Result {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for FindSubDistrictByParam")
	}

	var r0 <-chan helpers.Result
	if rf, ok := ret.Get(0).(func(context.Context, request.SubDistrict) <-chan helpers.Result); ok {
		r0 = rf(ctx, payload)
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