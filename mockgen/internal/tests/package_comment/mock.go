// Code generated by MockGen. DO NOT EDIT.
// Source: input.go
//
// Generated by this command:
//
//	mockgen -package empty_interface -destination mock.go -source input.go -write_package_comment=false
//

package empty_interface

import (
	gomock "go.uber.org/mock/gomock"
)

// MockEmpty is a mock of Empty interface.
type MockEmpty struct {
	ctrl     *gomock.Controller
	recorder *MockEmptyMockRecorder
	isgomock struct{}
}

// MockEmptyMockRecorder is the mock recorder for MockEmpty.
type MockEmptyMockRecorder struct {
	mock *MockEmpty
}

// NewMockEmpty creates a new mock instance.
func NewMockEmpty(ctrl *gomock.Controller) *MockEmpty {
	mock := &MockEmpty{ctrl: ctrl}
	mock.recorder = &MockEmptyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmpty) EXPECT() *MockEmptyMockRecorder {
	return m.recorder
}
