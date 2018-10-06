// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tonyalaribe/ninja/datalayer (interfaces: DataStore)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	datalayer "github.com/tonyalaribe/ninja/datalayer"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockDataStore) Connect(arg0 datalayer.DBConfig) (datalayer.DataStore, error) {
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(datalayer.DataStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect
func (mr *MockDataStoreMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDataStore)(nil).Connect), arg0)
}

// CreateCollection mocks base method
func (m *MockDataStore) CreateCollection(arg0 string, arg1, arg2 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "CreateCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCollection indicates an expected call of CreateCollection
func (mr *MockDataStoreMockRecorder) CreateCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCollection", reflect.TypeOf((*MockDataStore)(nil).CreateCollection), arg0, arg1, arg2)
}

// GetItem mocks base method
func (m *MockDataStore) GetItem(arg0, arg1 string) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetItem", arg0, arg1)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem
func (mr *MockDataStoreMockRecorder) GetItem(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockDataStore)(nil).GetItem), arg0, arg1)
}

// GetItems mocks base method
func (m *MockDataStore) GetItems(arg0 string, arg1 datalayer.QueryMeta) ([]map[string]interface{}, datalayer.ItemsResponseInfo, error) {
	ret := m.ctrl.Call(m, "GetItems", arg0, arg1)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(datalayer.ItemsResponseInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetItems indicates an expected call of GetItems
func (mr *MockDataStoreMockRecorder) GetItems(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockDataStore)(nil).GetItems), arg0, arg1)
}

// GetSchema mocks base method
func (m *MockDataStore) GetSchema(arg0 string) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetSchema", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (mr *MockDataStoreMockRecorder) GetSchema(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockDataStore)(nil).GetSchema), arg0)
}

// SaveItem mocks base method
func (m *MockDataStore) SaveItem(arg0, arg1 string, arg2 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "SaveItem", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveItem indicates an expected call of SaveItem
func (mr *MockDataStoreMockRecorder) SaveItem(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveItem", reflect.TypeOf((*MockDataStore)(nil).SaveItem), arg0, arg1, arg2)
}