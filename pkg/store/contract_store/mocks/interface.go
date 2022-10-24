// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/store/contract_store/interface.go

// Package mock_contractstore is a generated GoMock package.
package mock_contractstore

import (
	reflect "reflect"

	model "github.com/consolelabs/indexer-api/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIContract is a mock of IContract interface.
type MockIContract struct {
	ctrl     *gomock.Controller
	recorder *MockIContractMockRecorder
}

// MockIContractMockRecorder is the mock recorder for MockIContract.
type MockIContractMockRecorder struct {
	mock *MockIContract
}

// NewMockIContract creates a new mock instance.
func NewMockIContract(ctrl *gomock.Controller) *MockIContract {
	mock := &MockIContract{ctrl: ctrl}
	mock.recorder = &MockIContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIContract) EXPECT() *MockIContractMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIContract) Get() ([]*model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]*model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIContractMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIContract)(nil).Get))
}

// GetByAddress mocks base method.
func (m *MockIContract) GetByAddress(address string) (*model.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAddress", address)
	ret0, _ := ret[0].(*model.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAddress indicates an expected call of GetByAddress.
func (mr *MockIContractMockRecorder) GetByAddress(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAddress", reflect.TypeOf((*MockIContract)(nil).GetByAddress), address)
}

// Save mocks base method.
func (m *MockIContract) Save(new *model.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", new)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIContractMockRecorder) Save(new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIContract)(nil).Save), new)
}

// Update mocks base method.
func (m *MockIContract) Update(new *model.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", new)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIContractMockRecorder) Update(new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIContract)(nil).Update), new)
}
