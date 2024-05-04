// Code generated by mockery v2.42.3. DO NOT EDIT.

package repository

import (
	db "github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ObjectRepository is an autogenerated mock type for the ObjectRepository type
type ObjectRepository[T interface{ db.Object }] struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *ObjectRepository[T]) Delete(_a0 uuid.UUID) (*T, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*T, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *T); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *ObjectRepository[T]) Get(_a0 uuid.UUID) (*T, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*T, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *T); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: page, size
func (_m *ObjectRepository[T]) GetAll(page int, size int) ([]*T, error) {
	ret := _m.Called(page, size)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*T
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*T, error)); ok {
		return rf(page, size)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*T); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*T)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: record
func (_m *ObjectRepository[T]) Insert(record *T) (*T, error) {
	ret := _m.Called(record)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(*T) (*T, error)); ok {
		return rf(record)
	}
	if rf, ok := ret.Get(0).(func(*T) *T); ok {
		r0 = rf(record)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(*T) error); ok {
		r1 = rf(record)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: record
func (_m *ObjectRepository[T]) Update(record *T) (*T, error) {
	ret := _m.Called(record)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(*T) (*T, error)); ok {
		return rf(record)
	}
	if rf, ok := ret.Get(0).(func(*T) *T); ok {
		r0 = rf(record)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(*T) error); ok {
		r1 = rf(record)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewObjectRepository creates a new instance of ObjectRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObjectRepository[T interface{ db.Object }](t interface {
	mock.TestingT
	Cleanup(func())
}) *ObjectRepository[T] {
	mock := &ObjectRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
