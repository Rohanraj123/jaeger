// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	tracestore "github.com/jaegertracing/jaeger/storage_v2/tracestore"
	mock "github.com/stretchr/testify/mock"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *Factory) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTraceReader provides a mock function with no fields
func (_m *Factory) CreateTraceReader() (tracestore.Reader, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateTraceReader")
	}

	var r0 tracestore.Reader
	var r1 error
	if rf, ok := ret.Get(0).(func() (tracestore.Reader, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() tracestore.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tracestore.Reader)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTraceWriter provides a mock function with no fields
func (_m *Factory) CreateTraceWriter() (tracestore.Writer, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateTraceWriter")
	}

	var r0 tracestore.Writer
	var r1 error
	if rf, ok := ret.Get(0).(func() (tracestore.Writer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() tracestore.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tracestore.Writer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: ctx
func (_m *Factory) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFactory creates a new instance of Factory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *Factory {
	mock := &Factory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
