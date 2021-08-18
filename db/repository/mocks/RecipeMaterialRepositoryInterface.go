// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "go-resepee-api/entity"

	mock "github.com/stretchr/testify/mock"
)

// RecipeMaterialRepositoryInterface is an autogenerated mock type for the RecipeMaterialRepositoryInterface type
type RecipeMaterialRepositoryInterface struct {
	mock.Mock
}

// FindByRecipeID provides a mock function with given fields: recipeID
func (_m *RecipeMaterialRepositoryInterface) FindByRecipeID(recipeID int) ([]entity.RecipeMaterial, error) {
	ret := _m.Called(recipeID)

	var r0 []entity.RecipeMaterial
	if rf, ok := ret.Get(0).(func(int) []entity.RecipeMaterial); ok {
		r0 = rf(recipeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.RecipeMaterial)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(recipeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: recipeMaterial
func (_m *RecipeMaterialRepositoryInterface) Store(recipeMaterial *entity.RecipeMaterial) (entity.RecipeMaterial, error) {
	ret := _m.Called(recipeMaterial)

	var r0 entity.RecipeMaterial
	if rf, ok := ret.Get(0).(func(*entity.RecipeMaterial) entity.RecipeMaterial); ok {
		r0 = rf(recipeMaterial)
	} else {
		r0 = ret.Get(0).(entity.RecipeMaterial)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.RecipeMaterial) error); ok {
		r1 = rf(recipeMaterial)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}