// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mengelbart/moqtransport (interfaces: ControlMessageSender)
//
// Generated by this command:
//
//	mockgen -build_flags=-tags=gomock -package moqtransport -self_package github.com/mengelbart/moqtransport -destination mock_control_message_sender_test.go github.com/mengelbart/moqtransport ControlMessageSender
//

// Package moqtransport is a generated GoMock package.
package moqtransport

import (
	reflect "reflect"

	wire "github.com/mengelbart/moqtransport/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockControlMessageSender is a mock of ControlMessageSender interface.
type MockControlMessageSender struct {
	ctrl     *gomock.Controller
	recorder *MockControlMessageSenderMockRecorder
}

// MockControlMessageSenderMockRecorder is the mock recorder for MockControlMessageSender.
type MockControlMessageSenderMockRecorder struct {
	mock *MockControlMessageSender
}

// NewMockControlMessageSender creates a new mock instance.
func NewMockControlMessageSender(ctrl *gomock.Controller) *MockControlMessageSender {
	mock := &MockControlMessageSender{ctrl: ctrl}
	mock.recorder = &MockControlMessageSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControlMessageSender) EXPECT() *MockControlMessageSenderMockRecorder {
	return m.recorder
}

// close mocks base method.
func (m *MockControlMessageSender) close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "close")
}

// close indicates an expected call of close.
func (mr *MockControlMessageSenderMockRecorder) close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "close", reflect.TypeOf((*MockControlMessageSender)(nil).close))
}

// enqueue mocks base method.
func (m *MockControlMessageSender) enqueue(arg0 wire.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "enqueue", arg0)
}

// enqueue indicates an expected call of enqueue.
func (mr *MockControlMessageSenderMockRecorder) enqueue(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "enqueue", reflect.TypeOf((*MockControlMessageSender)(nil).enqueue), arg0)
}
