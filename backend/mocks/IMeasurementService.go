// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ppo/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IMeasurementService is an autogenerated mock type for the IMeasurementService type
type IMeasurementService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, measurement
func (_m *IMeasurementService) Create(ctx context.Context, measurement *domain.Measurement) error {
	ret := _m.Called(ctx, measurement)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Measurement) error); ok {
		r0 = rf(ctx, measurement)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *IMeasurementService) DeleteById(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *IMeasurementService) GetAll(ctx context.Context) ([]*domain.Measurement, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*domain.Measurement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Measurement, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Measurement); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Measurement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *IMeasurementService) GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.Measurement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Measurement, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Measurement); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Measurement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRecipeId provides a mock function with given fields: ctx, ingredientId, recipeId
func (_m *IMeasurementService) GetByRecipeId(ctx context.Context, ingredientId uuid.UUID, recipeId uuid.UUID) (*domain.Measurement, int, error) {
	ret := _m.Called(ctx, ingredientId, recipeId)

	if len(ret) == 0 {
		panic("no return value specified for GetByRecipeId")
	}

	var r0 *domain.Measurement
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) (*domain.Measurement, int, error)); ok {
		return rf(ctx, ingredientId, recipeId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) *domain.Measurement); ok {
		r0 = rf(ctx, ingredientId, recipeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Measurement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID) int); ok {
		r1 = rf(ctx, ingredientId, recipeId)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r2 = rf(ctx, ingredientId, recipeId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, measurement
func (_m *IMeasurementService) Update(ctx context.Context, measurement *domain.Measurement) error {
	ret := _m.Called(ctx, measurement)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Measurement) error); ok {
		r0 = rf(ctx, measurement)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLink provides a mock function with given fields: ctx, linkId, measurementId, amount
func (_m *IMeasurementService) UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error {
	ret := _m.Called(ctx, linkId, measurementId, amount)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLink")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, int) error); ok {
		r0 = rf(ctx, linkId, measurementId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIMeasurementService creates a new instance of IMeasurementService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMeasurementService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMeasurementService {
	mock := &IMeasurementService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}