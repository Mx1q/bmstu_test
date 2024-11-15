// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ppo/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ISaladTypeService is an autogenerated mock type for the ISaladTypeService type
type ISaladTypeService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, saladType
func (_m *ISaladTypeService) Create(ctx context.Context, saladType *domain.SaladType) error {
	ret := _m.Called(ctx, saladType)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SaladType) error); ok {
		r0 = rf(ctx, saladType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *ISaladTypeService) DeleteById(ctx context.Context, id uuid.UUID) error {
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

// GetAll provides a mock function with given fields: ctx, page
func (_m *ISaladTypeService) GetAll(ctx context.Context, page int) ([]*domain.SaladType, int, error) {
	ret := _m.Called(ctx, page)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*domain.SaladType
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]*domain.SaladType, int, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []*domain.SaladType); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.SaladType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) int); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int) error); ok {
		r2 = rf(ctx, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllBySaladId provides a mock function with given fields: ctx, saladId
func (_m *ISaladTypeService) GetAllBySaladId(ctx context.Context, saladId uuid.UUID) ([]*domain.SaladType, error) {
	ret := _m.Called(ctx, saladId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllBySaladId")
	}

	var r0 []*domain.SaladType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*domain.SaladType, error)); ok {
		return rf(ctx, saladId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*domain.SaladType); ok {
		r0 = rf(ctx, saladId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.SaladType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, saladId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *ISaladTypeService) GetById(ctx context.Context, id uuid.UUID) (*domain.SaladType, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.SaladType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.SaladType, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.SaladType); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SaladType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Link provides a mock function with given fields: ctx, saladId, saladTypeId
func (_m *ISaladTypeService) Link(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	ret := _m.Called(ctx, saladId, saladTypeId)

	if len(ret) == 0 {
		panic("no return value specified for Link")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, saladId, saladTypeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unlink provides a mock function with given fields: ctx, saladId, saladTypeId
func (_m *ISaladTypeService) Unlink(ctx context.Context, saladId uuid.UUID, saladTypeId uuid.UUID) error {
	ret := _m.Called(ctx, saladId, saladTypeId)

	if len(ret) == 0 {
		panic("no return value specified for Unlink")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, saladId, saladTypeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, measurement
func (_m *ISaladTypeService) Update(ctx context.Context, measurement *domain.SaladType) error {
	ret := _m.Called(ctx, measurement)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SaladType) error); ok {
		r0 = rf(ctx, measurement)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewISaladTypeService creates a new instance of ISaladTypeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewISaladTypeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ISaladTypeService {
	mock := &ISaladTypeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
