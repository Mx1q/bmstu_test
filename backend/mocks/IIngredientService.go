// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ppo/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// IIngredientService is an autogenerated mock type for the IIngredientService type
type IIngredientService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, salad
func (_m *IIngredientService) Create(ctx context.Context, salad *domain.Ingredient) error {
	ret := _m.Called(ctx, salad)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Ingredient) error); ok {
		r0 = rf(ctx, salad)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *IIngredientService) DeleteById(ctx context.Context, id uuid.UUID) error {
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
func (_m *IIngredientService) GetAll(ctx context.Context, page int) ([]*domain.Ingredient, int, error) {
	ret := _m.Called(ctx, page)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*domain.Ingredient
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]*domain.Ingredient, int, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []*domain.Ingredient); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Ingredient)
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

// GetAllByRecipeId provides a mock function with given fields: ctx, id
func (_m *IIngredientService) GetAllByRecipeId(ctx context.Context, id uuid.UUID) ([]*domain.Ingredient, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAllByRecipeId")
	}

	var r0 []*domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*domain.Ingredient, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*domain.Ingredient); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *IIngredientService) GetById(ctx context.Context, id uuid.UUID) (*domain.Ingredient, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *domain.Ingredient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Ingredient, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Ingredient); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ingredient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Link provides a mock function with given fields: ctx, recipeId, ingredientId
func (_m *IIngredientService) Link(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) (uuid.UUID, error) {
	ret := _m.Called(ctx, recipeId, ingredientId)

	if len(ret) == 0 {
		panic("no return value specified for Link")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) (uuid.UUID, error)); ok {
		return rf(ctx, recipeId, ingredientId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) uuid.UUID); ok {
		r0 = rf(ctx, recipeId, ingredientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(ctx, recipeId, ingredientId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlink provides a mock function with given fields: ctx, recipeId, ingredientId
func (_m *IIngredientService) Unlink(ctx context.Context, recipeId uuid.UUID, ingredientId uuid.UUID) error {
	ret := _m.Called(ctx, recipeId, ingredientId)

	if len(ret) == 0 {
		panic("no return value specified for Unlink")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, recipeId, ingredientId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, salad
func (_m *IIngredientService) Update(ctx context.Context, salad *domain.Ingredient) error {
	ret := _m.Called(ctx, salad)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Ingredient) error); ok {
		r0 = rf(ctx, salad)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIIngredientService creates a new instance of IIngredientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIIngredientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IIngredientService {
	mock := &IIngredientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
