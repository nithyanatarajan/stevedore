// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/http/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockClient) Get(url string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), url)
}

// Do mocks base method
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockClientMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClient)(nil).Do), req)
}