// Code generated by MockGen. DO NOT EDIT.
// Source: gomockhelper_test.go

// Package mock_gomockhelper_test is a generated GoMock package.
package mock_gomockhelper_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDoer1 is a mock of Doer1 interface
type MockDoer1 struct {
	ctrl     *gomock.Controller
	recorder *MockDoer1MockRecorder
}

// MockDoer1MockRecorder is the mock recorder for MockDoer1
type MockDoer1MockRecorder struct {
	mock *MockDoer1
}

// NewMockDoer1 creates a new mock instance
func NewMockDoer1(ctrl *gomock.Controller) *MockDoer1 {
	mock := &MockDoer1{ctrl: ctrl}
	mock.recorder = &MockDoer1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer1) EXPECT() *MockDoer1MockRecorder {
	return m.recorder
}

// Do1 mocks base method
func (m *MockDoer1) Do1() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do1")
}

// Do1 indicates an expected call of Do1
func (mr *MockDoer1MockRecorder) Do1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do1", reflect.TypeOf((*MockDoer1)(nil).Do1))
}

// Do2 mocks base method
func (m *MockDoer1) Do2() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do2")
}

// Do2 indicates an expected call of Do2
func (mr *MockDoer1MockRecorder) Do2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do2", reflect.TypeOf((*MockDoer1)(nil).Do2))
}

// Do3 mocks base method
func (m *MockDoer1) Do3() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do3")
}

// Do3 indicates an expected call of Do3
func (mr *MockDoer1MockRecorder) Do3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do3", reflect.TypeOf((*MockDoer1)(nil).Do3))
}

// MockDoer2 is a mock of Doer2 interface
type MockDoer2 struct {
	ctrl     *gomock.Controller
	recorder *MockDoer2MockRecorder
}

// MockDoer2MockRecorder is the mock recorder for MockDoer2
type MockDoer2MockRecorder struct {
	mock *MockDoer2
}

// NewMockDoer2 creates a new mock instance
func NewMockDoer2(ctrl *gomock.Controller) *MockDoer2 {
	mock := &MockDoer2{ctrl: ctrl}
	mock.recorder = &MockDoer2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer2) EXPECT() *MockDoer2MockRecorder {
	return m.recorder
}

// DoWithReturns1 mocks base method
func (m *MockDoer2) DoWithReturns1(a int, b string) (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoWithReturns1", a, b)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// DoWithReturns1 indicates an expected call of DoWithReturns1
func (mr *MockDoer2MockRecorder) DoWithReturns1(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoWithReturns1", reflect.TypeOf((*MockDoer2)(nil).DoWithReturns1), a, b)
}

// DoWithReturns2 mocks base method
func (m *MockDoer2) DoWithReturns2(a int, b string) (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoWithReturns2", a, b)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// DoWithReturns2 indicates an expected call of DoWithReturns2
func (mr *MockDoer2MockRecorder) DoWithReturns2(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoWithReturns2", reflect.TypeOf((*MockDoer2)(nil).DoWithReturns2), a, b)
}

// MockDoer3 is a mock of Doer3 interface
type MockDoer3 struct {
	ctrl     *gomock.Controller
	recorder *MockDoer3MockRecorder
}

// MockDoer3MockRecorder is the mock recorder for MockDoer3
type MockDoer3MockRecorder struct {
	mock *MockDoer3
}

// NewMockDoer3 creates a new mock instance
func NewMockDoer3(ctrl *gomock.Controller) *MockDoer3 {
	mock := &MockDoer3{ctrl: ctrl}
	mock.recorder = &MockDoer3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDoer3) EXPECT() *MockDoer3MockRecorder {
	return m.recorder
}

// Do1 mocks base method
func (m *MockDoer3) Do1(a int, b string) (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do1", a, b)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// Do1 indicates an expected call of Do1
func (mr *MockDoer3MockRecorder) Do1(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do1", reflect.TypeOf((*MockDoer3)(nil).Do1), a, b)
}

// Do2 mocks base method
func (m *MockDoer3) Do2(c bool, d []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do2", c, d)
}

// Do2 indicates an expected call of Do2
func (mr *MockDoer3MockRecorder) Do2(c, d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do2", reflect.TypeOf((*MockDoer3)(nil).Do2), c, d)
}
