// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/VanaraID/clean-architecture/models"

// ITodoRepository is an autogenerated mock type for the ITodoRepository type
type ITodoRepository struct {
	mock.Mock
}

// FetchAllTodo provides a mock function with given fields:
func (_m *ITodoRepository) FetchAllTodo() ([]models.Todo, error) {
	ret := _m.Called()

	var r0 []models.Todo
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchTodo provides a mock function with given fields: id
func (_m *ITodoRepository) FetchTodo(id int) (models.Todo, error) {
	ret := _m.Called(id)

	var r0 models.Todo
	if rf, ok := ret.Get(0).(func(int) models.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Todo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
