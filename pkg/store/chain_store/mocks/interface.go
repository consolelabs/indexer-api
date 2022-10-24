// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/store/chain_store/interface.go

// Package mock_chainstore is a generated GoMock package.
package mock_chainstore

import (
	reflect "reflect"

	request "github.com/consolelabs/indexer-api/pkg/request"
	model "github.com/consolelabs/indexer-api/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIChain is a mock of IChain interface.
type MockIChain struct {
	ctrl     *gomock.Controller
	recorder *MockIChainMockRecorder
}

// MockIChainMockRecorder is the mock recorder for MockIChain.
type MockIChainMockRecorder struct {
	mock *MockIChain
}

// NewMockIChain creates a new mock instance.
func NewMockIChain(ctrl *gomock.Controller) *MockIChain {
	mock := &MockIChain{ctrl: ctrl}
	mock.recorder = &MockIChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIChain) EXPECT() *MockIChainMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIChain) Get(query request.GetChainsQuery) ([]model.Chain, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", query)
	ret0, _ := ret[0].([]model.Chain)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockIChainMockRecorder) Get(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIChain)(nil).Get), query)
}
