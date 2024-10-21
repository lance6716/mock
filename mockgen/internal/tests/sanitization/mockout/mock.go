// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/mock/mockgen/internal/tests/sanitization (interfaces: AnyMock)
//
// Generated by this command:
//
//	mockgen -destination mockout/mock.go -package mockout . AnyMock
//

// Package mockout is a generated GoMock package.
package mockout

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	any0 "go.uber.org/mock/mockgen/internal/tests/sanitization/any"
)

// MockAnyMock is a mock of AnyMock interface.
type MockAnyMock struct {
	ctrl     *gomock.Controller
	recorder *MockAnyMockMockRecorder
	isgomock struct{}
}

// MockAnyMockMockRecorder is the mock recorder for MockAnyMock.
type MockAnyMockMockRecorder struct {
	mock *MockAnyMock
}

// NewMockAnyMock creates a new mock instance.
func NewMockAnyMock(ctrl *gomock.Controller) *MockAnyMock {
	mock := &MockAnyMock{ctrl: ctrl}
	mock.recorder = &MockAnyMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnyMock) EXPECT() *MockAnyMockMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockAnyMock) Do(a *any0.Any, b int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do", a, b)
}

// Do indicates an expected call of Do.
func (mr *MockAnyMockMockRecorder) Do(a, b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockAnyMock)(nil).Do), a, b)
}
