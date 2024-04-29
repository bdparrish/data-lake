// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	db "github.com/codingexplorations/data-lake/common/pkg/models/v1/db"

	mock "github.com/stretchr/testify/mock"
)

// IngestProcessor is an autogenerated mock type for the IngestProcessor type
type IngestProcessor struct {
	mock.Mock
}

// ProcessFile provides a mock function with given fields: fileName
func (_m *IngestProcessor) ProcessFile(fileName string) (*db.Object, error) {
	ret := _m.Called(fileName)

	if len(ret) == 0 {
		panic("no return value specified for ProcessFile")
	}

	var r0 *db.Object
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*db.Object, error)); ok {
		return rf(fileName)
	}
	if rf, ok := ret.Get(0).(func(string) *db.Object); ok {
		r0 = rf(fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Object)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessFolder provides a mock function with given fields: folder
func (_m *IngestProcessor) ProcessFolder(folder string) ([]*db.Object, error) {
	ret := _m.Called(folder)

	if len(ret) == 0 {
		panic("no return value specified for ProcessFolder")
	}

	var r0 []*db.Object
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*db.Object, error)); ok {
		return rf(folder)
	}
	if rf, ok := ret.Get(0).(func(string) []*db.Object); ok {
		r0 = rf(folder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.Object)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(folder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIngestProcessor creates a new instance of IngestProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIngestProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *IngestProcessor {
	mock := &IngestProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}