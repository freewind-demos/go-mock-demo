// Code generated by MockGen. DO NOT EDIT.
// Source: hello/hello.go

// Package mock_hello is a generated GoMock package.
package mock_hello

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTalker is a mock of Talker interface
type MockTalker struct {
	ctrl     *gomock.Controller
	recorder *MockTalkerMockRecorder
}

// MockTalkerMockRecorder is the mock recorder for MockTalker
type MockTalkerMockRecorder struct {
	mock *MockTalker
}

// NewMockTalker creates a new mock instance
func NewMockTalker(ctrl *gomock.Controller) *MockTalker {
	mock := &MockTalker{ctrl: ctrl}
	mock.recorder = &MockTalkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTalker) EXPECT() *MockTalkerMockRecorder {
	return m.recorder
}

// SayHello mocks base method
func (m *MockTalker) SayHello(word string) string {
	ret := m.ctrl.Call(m, "SayHello", word)
	ret0, _ := ret[0].(string)
	return ret0
}

// SayHello indicates an expected call of SayHello
func (mr *MockTalkerMockRecorder) SayHello(word interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockTalker)(nil).SayHello), word)
}
