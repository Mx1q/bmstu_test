// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ppo/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *IUserRepository) Create(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *IUserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
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
func (_m *IUserRepository) GetAll(ctx context.Context, page int) ([]*domain.User, error) {
	ret := _m.Called(ctx, page)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]*domain.User, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []*domain.User); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *IUserRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *IUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user
func (_m *IUserRepository) Update(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
