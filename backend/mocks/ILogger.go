// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ILogger is an autogenerated mock type for the ILogger type
type ILogger struct {
	mock.Mock
}

// Errorf provides a mock function with given fields: message, args
func (_m *ILogger) Errorf(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: message, args
func (_m *ILogger) Fatalf(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: message, args
func (_m *ILogger) Infof(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: message, args
func (_m *ILogger) Warnf(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// NewILogger creates a new instance of ILogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILogger {
	mock := &ILogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
