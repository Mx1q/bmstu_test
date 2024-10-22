// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ppo/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IRecipeStepInteractor is an autogenerated mock type for the IRecipeStepInteractor type
type IRecipeStepInteractor struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, recipeStep
func (_m *IRecipeStepInteractor) Create(ctx context.Context, recipeStep *domain.RecipeStep) error {
	ret := _m.Called(ctx, recipeStep)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.RecipeStep) error); ok {
		r0 = rf(ctx, recipeStep)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByRecipeID provides a mock function with given fields: ctx, recipeId
func (_m *IRecipeStepInteractor) DeleteAllByRecipeID(ctx context.Context, recipeId uuid.UUID) error {
	ret := _m.Called(ctx, recipeId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllByRecipeID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, recipeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *IRecipeStepInteractor) DeleteById(ctx context.Context, id uuid.UUID) error {
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

// GetAllByRecipeID provides a mock function with given fields: ctx, recipeId
func (_m *IRecipeStepInteractor) GetAllByRecipeID(ctx context.Context, recipeId uuid.UUID) ([]*domain.RecipeStep, error) {
	ret := _m.Called(ctx, recipeId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByRecipeID")
	}

	var r0 []*domain.RecipeStep
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*domain.RecipeStep, error)); ok {
		return rf(ctx, recipeId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*domain.RecipeStep); ok {
		r0 = rf(ctx, recipeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.RecipeStep)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, recipeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *IRecipeStepInteractor) GetById(ctx context.Context, id uuid.UUID) (*domain.RecipeStep, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.RecipeStep
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.RecipeStep, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.RecipeStep); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RecipeStep)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, recipeStep
func (_m *IRecipeStepInteractor) Update(ctx context.Context, recipeStep *domain.RecipeStep) error {
	ret := _m.Called(ctx, recipeStep)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.RecipeStep) error); ok {
		r0 = rf(ctx, recipeStep)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIRecipeStepInteractor creates a new instance of IRecipeStepInteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRecipeStepInteractor(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRecipeStepInteractor {
	mock := &IRecipeStepInteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
